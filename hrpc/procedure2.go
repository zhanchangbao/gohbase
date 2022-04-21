package hrpc

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/zhanchangbao/gohbase/pb"
)

// GetProcedures to represent a GetProcedures request
type GetProcedures struct {
	base
}

// NewGetProcedures creates a new GetProceduresStruct with default fields
func NewGetProcedures() *GetProcedures {
	return &GetProcedures{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (p *GetProcedures) Name() string {
	return "GetProcedures"
}

// Description returns the description of this RPC call.
func (p *GetProcedures) Description() string {
	return p.Name()
}

// ToProto returns the Protobuf message to be sent
func (p *GetProcedures) ToProto() proto.Message {
	return &pb.GetProceduresRequest{}
}

// NewResponse returns the empty protobuf response
func (p *GetProcedures) NewResponse() proto.Message {
	return &pb.GetProceduresResponse{}
}

// ClusterProcedures to represent a cluster procedures request
type ClusterProcedures struct {
	base
}

// NewClusterProcedures creates a new ClusterStatusStruct with default fields
func NewClusterProcedures() *ClusterProcedures {
	return &ClusterProcedures{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (c *ClusterProcedures) Name() string {
	return "GetProcedures"
}

// Description returns the description of this RPC call.
func (c *ClusterProcedures) Description() string {
	return c.Name()
}

// ToProto returns the Protobuf message to be sent
func (c *ClusterProcedures) ToProto() proto.Message {
	return &pb.GetProceduresRequest{}
}

// NewResponse returns the empty protobuf response
func (c *ClusterProcedures) NewResponse() proto.Message {
	return &pb.GetProceduresResponse{}
}
