// Copyright (C) 2017  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package region

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/zhanchangbao/gohbase/hrpc"
	"github.com/zhanchangbao/gohbase/pb"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

var multiPool = sync.Pool{
	New: func() interface{} {
		return &multi{}
	},
}

func freeMulti(m *multi) {
	m.calls = m.calls[:0]
	m.regions = m.regions[:0]
	m.size = 0
	multiPool.Put(m)
}

type multi struct {
	ctxM sync.Mutex
	ctx  context.Context

	size  int
	calls []hrpc.Call
	// regions preserves the order of regions to match against RegionActionResults
	regions []hrpc.RegionInfo
}

func newMulti(ctx context.Context, queueSize int) *multi {
	m := multiPool.Get().(*multi)
	m.ctx = ctx
	m.size = queueSize
	return m
}

// Name returns the name of this RPC call.
func (m *multi) Name() string {
	return "Multi"
}

// Description returns the description of this RPC call.
func (m *multi) Description() string {
	return m.Name()
}

// ToProto converts all request in multi batch to a protobuf message.
func (m *multi) ToProto() proto.Message {
	msg, _, _ := m.toProto(false)
	return msg
}

type actions struct {
	pbs        []*pb.Action
	cellblocks [][]byte
}

func (m *multi) toProto(isCellblocks bool) (proto.Message, [][]byte, uint32) {
	// aggregate calls per region
	actionsPerReg := map[hrpc.RegionInfo]actions{}
	var size uint32

	for i, c := range m.calls {
		select {
		case <-c.Context().Done():
			// context has expired, don't bother sending it
			m.calls[i] = nil
			continue
		default:
		}

		var msg proto.Message
		var cellblocks [][]byte
		if s, ok := c.(canSerializeCellBlocks); isCellblocks && ok && s.CellBlocksEnabled() {
			var sz uint32
			msg, cellblocks, sz = s.SerializeCellBlocks()
			size += sz
		} else {
			msg = c.ToProto()
		}

		a := &pb.Action{
			Index: proto.Uint32(uint32(i) + 1), // +1 because 0 index means there's no index
		}

		switch r := msg.(type) {
		case *pb.GetRequest:
			a.Get = r.Get
		case *pb.MutateRequest:
			a.Mutation = r.Mutation
		default:
			panic(fmt.Sprintf("unsupported call type for Multi: %T", c))
		}

		as, ok := actionsPerReg[c.Region()]
		if !ok {
			as = actions{}
		}
		as.pbs = append(as.pbs, a)
		as.cellblocks = append(as.cellblocks, cellblocks...)

		actionsPerReg[c.Region()] = as
	}

	// construct the multi proto
	ra := make([]*pb.RegionAction, len(actionsPerReg))
	m.regions = make([]hrpc.RegionInfo, len(actionsPerReg))

	var cellblocks [][]byte
	i := 0
	for r, as := range actionsPerReg {
		ra[i] = &pb.RegionAction{
			Region: &pb.RegionSpecifier{
				Type:  hrpc.RegionSpecifierRegionName,
				Value: r.Name(),
			},
			Action: as.pbs,
		}
		cellblocks = append(cellblocks, as.cellblocks...)
		// Track the order of RegionActions,
		// so that we can handle whole region exceptions.
		m.regions[i] = r
		i++
	}
	return &pb.MultiRequest{RegionAction: ra}, cellblocks, size
}

func (m *multi) SerializeCellBlocks() (proto.Message, [][]byte, uint32) {
	return m.toProto(true)
}

func (m *multi) CellBlocksEnabled() bool {
	// TODO: maybe have some global client option
	return true
}

// NewResponse creates an empty protobuf message to read the response of this RPC.
func (m *multi) NewResponse() proto.Message {
	return &pb.MultiResponse{}
}

// DeserializeCellBlocks deserializes action results from cell blocks.
func (m *multi) DeserializeCellBlocks(msg proto.Message, b []byte) (uint32, error) {
	mr := msg.(*pb.MultiResponse)

	var nread uint32
	for _, rar := range mr.GetRegionActionResult() {
		if e := rar.GetException(); e != nil {
			if l := len(rar.GetResultOrException()); l != 0 {
				return 0, fmt.Errorf(
					"got exception for region, but still have %d result(s) returned from it", l)
			}
			continue
		}

		for _, roe := range rar.GetResultOrException() {
			e := roe.GetException()
			r := roe.GetResult()
			i := roe.GetIndex()

			if i == 0 {
				return 0, errors.New("no index for result in multi response")
			} else if r == nil && e == nil {
				return 0, errors.New("no result or exception for action in multi response")
			} else if r != nil && e != nil {
				return 0, errors.New("got result and exception for action in multi response")
			} else if e != nil {
				continue
			}

			c := m.get(i)                     // TODO: maybe return error if it's out-of-bounds
			d := c.(canDeserializeCellBlocks) // let it panic, because then it's our bug

			response := c.NewResponse()
			switch rsp := response.(type) {
			case *pb.GetResponse:
				rsp.Result = r
			case *pb.MutateResponse:
				rsp.Result = r
			default:
				panic(fmt.Sprintf("unsupported response type for Multi: %T", response))
			}

			// TODO: don't bother deserializing if the call's context has already expired
			n, err := d.DeserializeCellBlocks(response, b[nread:])
			if err != nil {
				return 0, fmt.Errorf(
					"error deserializing cellblocks for %q call as part of MultiResponse: %v",
					c.Name(), err)
			}
			nread += n
		}
	}
	return nread, nil
}

func (m *multi) returnResults(msg proto.Message, err error) {
	defer freeMulti(m)

	if err != nil {
		for _, c := range m.calls {
			if c == nil {
				continue
			}
			c.ResultChan() <- hrpc.RPCResult{Error: err}
		}
		return
	}

	mr := msg.(*pb.MultiResponse)

	// Here we can assume that everything has been deserialized correctly.
	// Dispatch results to appropriate calls.
	for i, rar := range mr.GetRegionActionResult() {
		if e := rar.GetException(); e != nil {
			// Got an exception for the whole region,
			// fail all the calls for that region.
			reg := m.regions[i]

			err := exceptionToError(*e.Name, string(e.Value))
			for _, c := range m.calls {
				if c == nil {
					continue
				}
				if c.Region() == reg {
					c.ResultChan() <- hrpc.RPCResult{Error: err}
				}
			}
			continue
		}

		for _, roe := range rar.GetResultOrException() {
			i := roe.GetIndex()
			e := roe.GetException()
			r := roe.GetResult()

			c := m.get(i)

			// TODO: don't bother if the call's context has already expired

			if e != nil {
				c.ResultChan() <- hrpc.RPCResult{
					Error: exceptionToError(*e.Name, string(e.Value)),
				}
				continue
			}

			response := c.NewResponse()
			switch rsp := response.(type) {
			case *pb.GetResponse:
				rsp.Result = r
			case *pb.MutateResponse:
				rsp.Result = r
			default:
				panic(fmt.Sprintf("unsupported response type for Multi: %T", response))
			}

			c.ResultChan() <- hrpc.RPCResult{Msg: response}
		}
	}
}

// add adds the call and returns wether the batch is full.
func (m *multi) add(call hrpc.Call) bool {
	msp := trace.SpanFromContext(m.Context())
	if msp.IsRecording() {
		csp := trace.SpanContextFromContext(call.Context())
		if csp.HasTraceID() {
			traceID := csp.TraceID().String()
			if !csp.IsSampled() {
				traceID += " (not sampled)"
			}
			msp.AddEvent("enqueue", trace.WithAttributes(
				attribute.String("traceid", traceID),
			))
		}
	}

	m.calls = append(m.calls, call)
	return len(m.calls) == m.size
}

// len returns number of batched calls.
func (m *multi) len() int {
	return len(m.calls)
}

// get retruns an rpc at index. Indicies start from 1 since 0 means that
// region server didn't set an index for the action result.
func (m *multi) get(i uint32) hrpc.Call {
	if i == 0 {
		panic("index cannot be 0")
	}
	return m.calls[i-1]
}

// Table is not supported for Multi.
func (m *multi) Table() []byte {
	panic("'Table' is not supported for 'Multi'")
}

// Reqion is not supported for Multi.
func (m *multi) Region() hrpc.RegionInfo {
	panic("'Region' is not supported for 'Multi'")
}

// SetRegion is not supported for Multi.
func (m *multi) SetRegion(r hrpc.RegionInfo) {
	panic("'SetRegion' is not supported for 'Multi'")
}

// ResultChan is not supported for Multi.
func (m *multi) ResultChan() chan hrpc.RPCResult {
	panic("'ResultChan' is not supported for 'Multi'")
}

// Context is not supported for Multi.
func (m *multi) Context() context.Context {
	// TODO: maybe pick the one with the longest deadline and use a context that has that deadline?
	m.ctxM.Lock()
	defer m.ctxM.Unlock()

	return m.ctx
}

// String returns a description of this call
func (m *multi) String() string {
	return "MULTI"
}

// SetContext is used for tracing implementations
func (m *multi) SetContext(ctx context.Context) {
	m.ctxM.Lock()
	defer m.ctxM.Unlock()

	m.ctx = ctx
}

// Key is not supported for Multi RPC.
func (m *multi) Key() []byte {
	panic("'Key' is not supported for 'Multi'")
}

// callCount returns how many calls are represented
func (m *multi) callCount() int {
	return len(m.calls)
}

// addSendEventsToCallSpans adds a send event
// to each of the calls represented. This
// will run when they are actually sent
func (m *multi) addSendEventsToCallSpans() {
	spCtx := trace.SpanContextFromContext(m.Context())
	traceID := ""

	if spCtx.HasTraceID() {
		traceID = spCtx.TraceID().String()
		if !spCtx.IsSampled() {
			traceID += " (not sampled)"
		}
	}

	for _, c := range m.calls {
		sp := trace.SpanFromContext(c.Context())
		sp.AddEvent(
			"send",
			trace.WithAttributes(attribute.String("traceID", traceID)),
		)
	}
}
