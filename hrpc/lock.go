package hrpc

import (
	"context"

	"github.com/zhanchangbao/gohbase/pb"
	"google.golang.org/protobuf/proto"
)

// Lock to represent a cluster Lock request
type Lock struct {
	base
}

// NewLock creates a new LockStruct with default fields
func NewLock() *Lock {
	return &Lock{
		base{
			ctx:      context.Background(),
			table:    []byte{},
			resultch: make(chan RPCResult, 1),
		},
	}
}

// Name returns the name of the rpc function
func (c *Lock) Name() string {
	return "GetLocks"
}

// Description returns the description of this RPC call.
func (c *Lock) Description() string {
	return c.Name()
}

// ToProto returns the Protobuf message to be sent
func (c *Lock) ToProto() proto.Message {
	return &pb.GetLocksRequest{}
}

// NewResponse returns the empty protobuf response
func (c *Lock) NewResponse() proto.Message {
	return &pb.GetLocksResponse{}
}
