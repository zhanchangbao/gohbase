package hrpc

import (
	"context"

	"github.com/zhanchangbao/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// Procedures to represent a cluster procedures request
type Procedures struct {
	base
}

// NewProcedures creates a new ProceduresStruct with default fields
func NewProcedures() *Procedures {
	return &Procedures{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (c *Procedures) Name() string {
	return "GetProcedures"
}

// Description returns the description of this RPC call.
func (c *Procedures) Description() string {
	return c.Name()
}

// ToProto returns the Protobuf message to be sent
func (c *Procedures) ToProto() proto.Message {
	return &pb.GetProceduresRequest{}
}

// NewResponse returns the empty protobuf response
func (c *Procedures) NewResponse() proto.Message {
	return &pb.GetProceduresResponse{}
}
