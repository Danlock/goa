// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter GRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package server

import (
	"context"
	"strings"

	goa "goa.design/goa"
	chattersvc "goa.design/goa/examples/streaming/gen/chatter"
	chatterpb "goa.design/goa/examples/streaming/gen/grpc/chatter"
	"google.golang.org/grpc/metadata"
)

// EncodeLoginResponse encodes responses from the chatter login endpoint.
func EncodeLoginResponse(ctx context.Context, v interface{}) *chatterpb.LoginResponse {
	res := v.(string)
	resp := NewLoginResponse(res)
	return resp
}

// DecodeLoginRequest decodes requests sent to chatter login endpoint.
func DecodeLoginRequest(ctx context.Context, message *chatterpb.LoginRequest) (*chattersvc.LoginPayload, error) {
	var (
		payload *chattersvc.LoginPayload
		err     error
	)
	{
		var (
			user     string
			password string
		)
		{
			md, ok := metadata.FromIncomingContext(ctx)
			if ok {
				if v := md.Get("user"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("user", "metadata"))
				} else {
					user = v[0]
				}
				if v := md.Get("password"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("password", "metadata"))
				} else {
					password = v[0]
				}
			}
		}
		payload = NewLoginPayload(message, user, password)
	}
	return payload, err
}

// DecodeEchoerRequest decodes requests sent to chatter echoer endpoint.
func DecodeEchoerRequest(stream chatterpb.Chatter_EchoerServer) (*chattersvc.EchoerPayload, error) {
	var (
		payload *chattersvc.EchoerPayload
		err     error
	)
	{
		var (
			token string
		)
		{
			md, ok := metadata.FromIncomingContext(stream.Context())
			if ok {
				if v := md.Get("authorization"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
				} else {
					token = v[0]
				}
			}
		}
		payload = NewEchoerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}

// DecodeListenerRequest decodes requests sent to chatter listener endpoint.
func DecodeListenerRequest(stream chatterpb.Chatter_ListenerServer) (*chattersvc.ListenerPayload, error) {
	var (
		payload *chattersvc.ListenerPayload
		err     error
	)
	{
		var (
			token string
		)
		{
			md, ok := metadata.FromIncomingContext(stream.Context())
			if ok {
				if v := md.Get("authorization"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
				} else {
					token = v[0]
				}
			}
		}
		payload = NewListenerPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}

// DecodeSummaryRequest decodes requests sent to chatter summary endpoint.
func DecodeSummaryRequest(stream chatterpb.Chatter_SummaryServer) (*chattersvc.SummaryPayload, error) {
	var (
		payload *chattersvc.SummaryPayload
		err     error
	)
	{
		var (
			token string
		)
		{
			md, ok := metadata.FromIncomingContext(stream.Context())
			if ok {
				if v := md.Get("authorization"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
				} else {
					token = v[0]
				}
			}
		}
		payload = NewSummaryPayload(token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}

// DecodeHistoryRequest decodes requests sent to chatter history endpoint.
func DecodeHistoryRequest(stream chatterpb.Chatter_HistoryServer, message *chatterpb.HistoryRequest) (*chattersvc.HistoryPayload, error) {
	var (
		payload *chattersvc.HistoryPayload
		err     error
	)
	{
		var (
			view  string
			token string
		)
		{
			md, ok := metadata.FromIncomingContext(stream.Context())
			if ok {
				if v := md.Get("view"); len(v) > 0 {
					view = v[0]
				}
				if v := md.Get("authorization"); len(v) == 0 {
					err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
				} else {
					token = v[0]
				}
			}
		}
		payload = NewHistoryPayload(message, view, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}
