// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// secured_service GRPC server encoders and decoders
//
// Command:
// $ goa gen goa.design/goa/examples/security/design -o
// $(GOPATH)/src/goa.design/goa/examples/security

package server

import (
	"context"
	"strings"

	goa "goa.design/goa"
	"goa.design/goa/examples/security/gen/grpc/secured_service/pb"
	securedservice "goa.design/goa/examples/security/gen/secured_service"
	goagrpc "goa.design/goa/grpc"
	"google.golang.org/grpc/metadata"
)

// EncodeSigninResponse encodes responses from the "secured_service" service
// "signin" endpoint.
func EncodeSigninResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(*securedservice.Creds)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*securedservice.Creds", v)
	}
	resp := NewSigninResponse(result)
	return resp, nil
}

// DecodeSigninRequest decodes requests sent to "secured_service" service
// "signin" endpoint.
func DecodeSigninRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		username string
		password string
		err      error
	)
	{
		if vals := md.Get("username"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("username", "metadata"))
		} else {
			username = vals[0]
		}
		if vals := md.Get("password"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("password", "metadata"))
		} else {
			password = vals[0]
		}
	}
	var (
		message *pb.SigninRequest
		ok      bool
	)
	{
		if message, ok = v.(*pb.SigninRequest); !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "signin", "*pb.SigninRequest", v)
		}
	}
	var (
		payload *securedservice.SigninPayload
	)
	{
		payload = NewSigninPayload(message, username, password)
	}
	return payload, err
}

// EncodeSecureResponse encodes responses from the "secured_service" service
// "secure" endpoint.
func EncodeSecureResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "secure", "string", v)
	}
	resp := NewSecureResponse(result)
	return resp, nil
}

// DecodeSecureRequest decodes requests sent to "secured_service" service
// "secure" endpoint.
func DecodeSecureRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	var (
		message *pb.SecureRequest
		ok      bool
	)
	{
		if message, ok = v.(*pb.SecureRequest); !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "secure", "*pb.SecureRequest", v)
		}
	}
	var (
		payload *securedservice.SecurePayload
	)
	{
		payload = NewSecurePayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}

// EncodeDoublySecureResponse encodes responses from the "secured_service"
// service "doubly_secure" endpoint.
func EncodeDoublySecureResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "doubly_secure", "string", v)
	}
	resp := NewDoublySecureResponse(result)
	return resp, nil
}

// DecodeDoublySecureRequest decodes requests sent to "secured_service" service
// "doubly_secure" endpoint.
func DecodeDoublySecureRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		token string
		err   error
	)
	{
		if vals := md.Get("authorization"); len(vals) == 0 {
			err = goa.MergeErrors(err, goa.MissingFieldError("authorization", "metadata"))
		} else {
			token = vals[0]
		}
	}
	var (
		message *pb.DoublySecureRequest
		ok      bool
	)
	{
		if message, ok = v.(*pb.DoublySecureRequest); !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "doubly_secure", "*pb.DoublySecureRequest", v)
		}
	}
	var (
		payload *securedservice.DoublySecurePayload
	)
	{
		payload = NewDoublySecurePayload(message, token)
		if strings.Contains(payload.Token, " ") {
			// Remove authorization scheme prefix (e.g. "Bearer")
			cred := strings.SplitN(payload.Token, " ", 2)[1]
			payload.Token = cred
		}
	}
	return payload, err
}

// EncodeAlsoDoublySecureResponse encodes responses from the "secured_service"
// service "also_doubly_secure" endpoint.
func EncodeAlsoDoublySecureResponse(ctx context.Context, v interface{}, hdr, trlr *metadata.MD) (interface{}, error) {
	result, ok := v.(string)
	if !ok {
		return nil, goagrpc.ErrInvalidType("secured_service", "also_doubly_secure", "string", v)
	}
	resp := NewAlsoDoublySecureResponse(result)
	return resp, nil
}

// DecodeAlsoDoublySecureRequest decodes requests sent to "secured_service"
// service "also_doubly_secure" endpoint.
func DecodeAlsoDoublySecureRequest(ctx context.Context, v interface{}, md metadata.MD) (interface{}, error) {
	var (
		oauthToken string
		token      string
		err        error
	)
	{
		if vals := md.Get("oauth"); len(vals) > 0 {
			oauthToken = vals[0]
		}
		if vals := md.Get("authorization"); len(vals) > 0 {
			token = vals[0]
		}
	}
	var (
		message *pb.AlsoDoublySecureRequest
		ok      bool
	)
	{
		if message, ok = v.(*pb.AlsoDoublySecureRequest); !ok {
			return nil, goagrpc.ErrInvalidType("secured_service", "also_doubly_secure", "*pb.AlsoDoublySecureRequest", v)
		}
	}
	var (
		payload *securedservice.AlsoDoublySecurePayload
	)
	{
		payload = NewAlsoDoublySecurePayload(message, oauthToken, token)
		if payload.Token != nil {
			if strings.Contains(*payload.Token, " ") {
				// Remove authorization scheme prefix (e.g. "Bearer")
				cred := strings.SplitN(*payload.Token, " ", 2)[1]
				payload.Token = &cred
			}
		}
	}
	return payload, err
}
