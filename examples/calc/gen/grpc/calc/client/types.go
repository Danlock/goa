// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// calc gRPC client types
//
// Command:
// $ goa gen goa.design/goa/examples/calc/design -o
// $(GOPATH)/src/goa.design/goa/examples/calc

package client

import (
	calcsvc "goa.design/goa/examples/calc/gen/calc"
	"goa.design/goa/examples/calc/gen/grpc/calc/pb"
)

// NewAddRequest builds the gRPC request type from the payload of the "add"
// endpoint of the "calc" service.
func NewAddRequest(p *calcsvc.AddPayload) *pb.AddRequest {
	v := &pb.AddRequest{
		A: int32(p.A),
		B: int32(p.B),
	}
	return v
}

// NewAddResponse builds the result type of the "add" endpoint of the "calc"
// service from the gRPC response type.
func NewAddResponse(resp *pb.AddResponse) int {
	v := int(resp.Field)
	return v
}
