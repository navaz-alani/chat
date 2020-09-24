package service

import "github.com/navaz-alani/chat/pkg/payload"

type Client interface {
	// the Send method returns a channel over which the caller can send messages
	// to the client.
	Send() chan<- payload.Payload
	// ID returns a string which identifies the client uniquely.
	ID() string
	// Listen fires off the read/write routines for the associated client.
	Listen()
}
