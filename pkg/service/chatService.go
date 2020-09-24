package service

import "github.com/navaz-alani/chat/pkg/payload"

type ChatService interface {
	AddParticipant(c Client) error
	RemoveParticipant(c Client) error
	Distribute() chan<- payload.Payload
	Serve()
}
