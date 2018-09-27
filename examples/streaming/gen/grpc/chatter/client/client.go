// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter GRPC client
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package client

import (
	"context"

	goa "goa.design/goa"
	chattersvc "goa.design/goa/examples/streaming/gen/chatter"
	chatterpb "goa.design/goa/examples/streaming/gen/grpc/chatter"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli chatterpb.ChatterClient
	opts    []grpc.CallOption
}

// echoerClientStream implements the chattersvc.EchoerClientStream.%!s(MISSING)
// interface.
type echoerClientStream struct {
	stream chatterpb.Chatter_EchoerClient
}

// listenerClientStream implements the
// chattersvc.ListenerClientStream.%!s(MISSING) interface.
type listenerClientStream struct {
	stream chatterpb.Chatter_ListenerClient
}

// summaryClientStream implements the
// chattersvc.SummaryClientStream.%!s(MISSING) interface.
type summaryClientStream struct {
	stream chatterpb.Chatter_SummaryClient
}

// historyClientStream implements the
// chattersvc.HistoryClientStream.%!s(MISSING) interface.
type historyClientStream struct {
	stream chatterpb.Chatter_HistoryClient
}

// NewClient instantiates gRPC client for all the chatter service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: chatterpb.NewChatterClient(cc),
		opts:    opts,
	}
}

// Login calls the "Login" function in chatterpb.ChatterClient interface.
func (c *Client) Login() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*chattersvc.LoginPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("chatter", "login", "*chattersvc.LoginPayload", v)
		}
		ctx, req := EncodeLoginRequest(ctx, p)
		resp, err := c.grpccli.Login(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		return DecodeLoginResponse(ctx, resp)
	}
}

// Echoer calls the "Echoer" function in chatterpb.ChatterClient interface.
func (c *Client) Echoer() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*chattersvc.EchoerPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("chatter", "echoer", "*chattersvc.EchoerPayload", v)
		}
		ctx = EncodeEchoerRequest(ctx, p)
		stream, err := c.grpccli.Echoer(ctx, c.opts...)
		if err != nil {
			return nil, err
		}
		return &echoerClientStream{stream: stream}, nil
	}
}

// Listener calls the "Listener" function in chatterpb.ChatterClient interface.
func (c *Client) Listener() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*chattersvc.ListenerPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("chatter", "listener", "*chattersvc.ListenerPayload", v)
		}
		ctx = EncodeListenerRequest(ctx, p)
		stream, err := c.grpccli.Listener(ctx, c.opts...)
		if err != nil {
			return nil, err
		}
		return &listenerClientStream{stream: stream}, nil
	}
}

// Summary calls the "Summary" function in chatterpb.ChatterClient interface.
func (c *Client) Summary() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*chattersvc.SummaryPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("chatter", "summary", "*chattersvc.SummaryPayload", v)
		}
		ctx = EncodeSummaryRequest(ctx, p)
		stream, err := c.grpccli.Summary(ctx, c.opts...)
		if err != nil {
			return nil, err
		}
		return &summaryClientStream{stream: stream}, nil
	}
}

// History calls the "History" function in chatterpb.ChatterClient interface.
func (c *Client) History() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		p, ok := v.(*chattersvc.HistoryPayload)
		if !ok {
			return nil, goagrpc.ErrInvalidType("chatter", "history", "*chattersvc.HistoryPayload", v)
		}
		ctx, req := EncodeHistoryRequest(ctx, p)
		stream, err := c.grpccli.History(ctx, req, c.opts...)
		if err != nil {
			return nil, err
		}
		return &historyClientStream{stream: stream}, nil
	}
}

// Recv reads instances of "chatterpb.EchoerResponse" from the "echoer"
// endpoint gRPC stream.
func (s *echoerClientStream) Recv() (string, error) {
	var res string
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	res = NewEchoerResponse(v)
	return res, nil
}

// Send streams instances of "string" to the "echoer" endpoint gRPC stream.
func (s *echoerClientStream) Send(res string) error {
	v := NewEchoerStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *echoerClientStream) Close() error {
	// nothing to do here
	return nil
}

// Send streams instances of "string" to the "listener" endpoint gRPC stream.
func (s *listenerClientStream) Send(res string) error {
	v := NewListenerStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *listenerClientStream) Close() error {
	// nothing to do here
	return nil
}

// CloseAndRecv reads instances of "chatterpb.ChatSummaryCollection" from the
// "summary" endpoint gRPC stream.
func (s *summaryClientStream) CloseAndRecv() (chattersvc.ChatSummaryCollection, error) {
	var res chattersvc.ChatSummaryCollection
	v, err := s.stream.CloseAndRecv()
	if err != nil {
		return res, err
	}
	res = NewChatSummaryCollection(v)
	return res, nil
}

// Send streams instances of "string" to the "summary" endpoint gRPC stream.
func (s *summaryClientStream) Send(res string) error {
	v := NewSummaryStreamingRequest(res)
	return s.stream.Send(v)
}

// Recv reads instances of "chatterpb.HistoryResponse" from the "history"
// endpoint gRPC stream.
func (s *historyClientStream) Recv() (*chattersvc.ChatSummary, error) {
	var res *chattersvc.ChatSummary
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	res = NewChatSummary(v)
	return res, nil
}

func (s *historyClientStream) Close() error {
	// nothing to do here
	return nil
}

// SetView sets the view.
func (s *historyClientStream) SetView(view string) {
}
