// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// chatter gRPC client CLI support package
//
// Command:
// $ goa gen goa.design/goa/examples/streaming/design -o
// $(GOPATH)/src/goa.design/goa/examples/streaming

package client

import (
	chattersvc "goa.design/goa/examples/streaming/gen/chatter"
)

// BuildLoginPayload builds the payload for the chatter login endpoint from CLI
// flags.
func BuildLoginPayload(chatterLoginUser string, chatterLoginPassword string) (*chattersvc.LoginPayload, error) {
	var user string
	{
		user = chatterLoginUser
	}
	var password string
	{
		password = chatterLoginPassword
	}
	payload := &chattersvc.LoginPayload{}
	payload.User = user
	payload.Password = password
	return payload, nil
}

// BuildEchoerPayload builds the payload for the chatter echoer endpoint from
// CLI flags.
func BuildEchoerPayload(chatterEchoerToken string) (*chattersvc.EchoerPayload, error) {
	var token string
	{
		token = chatterEchoerToken
	}
	payload := &chattersvc.EchoerPayload{}
	payload.Token = token
	return payload, nil
}

// BuildListenerPayload builds the payload for the chatter listener endpoint
// from CLI flags.
func BuildListenerPayload(chatterListenerToken string) (*chattersvc.ListenerPayload, error) {
	var token string
	{
		token = chatterListenerToken
	}
	payload := &chattersvc.ListenerPayload{}
	payload.Token = token
	return payload, nil
}

// BuildSummaryPayload builds the payload for the chatter summary endpoint from
// CLI flags.
func BuildSummaryPayload(chatterSummaryToken string) (*chattersvc.SummaryPayload, error) {
	var token string
	{
		token = chatterSummaryToken
	}
	payload := &chattersvc.SummaryPayload{}
	payload.Token = token
	return payload, nil
}

// BuildHistoryPayload builds the payload for the chatter history endpoint from
// CLI flags.
func BuildHistoryPayload(chatterHistoryView string, chatterHistoryToken string) (*chattersvc.HistoryPayload, error) {
	var view *string
	{
		if chatterHistoryView != "" {
			view = &chatterHistoryView
		}
	}
	var token string
	{
		token = chatterHistoryToken
	}
	payload := &chattersvc.HistoryPayload{}
	payload.View = view
	payload.Token = token
	return payload, nil
}
