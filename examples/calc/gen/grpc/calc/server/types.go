// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc gRPC server types
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package server

import (
	calcsvc "goa.design/goa/examples/calc/gen/calc"
	"goa.design/goa/examples/calc/gen/grpc/calc/pb"
)

// NewAddPayload builds the payload of the "add" endpoint of the "calc" service
// from the gRPC request type.
func NewAddPayload(message *pb.AddRequest) *calcsvc.AddPayload {
	v := &calcsvc.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}
	return v
}

// NewAddResponse builds the gRPC response type from the result of the "add"
// endpoint of the "calc" service.
func NewAddResponse(res int) *pb.AddResponse {
	v := &pb.AddResponse{}
	v.Field = int32(res)
	return v
}
