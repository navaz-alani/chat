package service

import (
	"fmt"

	"github.com/navaz-alani/chat/pkg/payload"
)

type Service struct {
	connected map[string]Client
	dist      chan payload.Payload
}

func NewService() (*Service, error) {
	return &Service{
		connected: make(map[string]Client),
		dist:      make(chan payload.Payload),
	}, nil
}

func (s *Service) AddParticipant(c Client) error {
	// if already connected, do nothing
	if _, ok := s.connected[c.ID()]; !ok {
		s.connected[c.ID()] = c
		c.Listen()
		return nil
	}
	// return error for multiple connections
	return fmt.Errorf("particpant is already connected")
}

func (s *Service) RemoveParticipant(c Client) error {
	// if client doesn't exist, do nothing
	if _, ok := s.connected[c.ID()]; ok {
		delete(s.connected, c.ID())
	}
	return nil
}

func (s *Service) Distribute() chan<- payload.Payload {
	return s.dist
}

func (s *Service) Serve() {
	go s.distribute()
}

// distribution routine
func (s *Service) distribute() {
	for {
		select {
		case p := <-s.dist:
			{
				// check if sender exists. If not, message contents could be corrupted
				sender, ok := s.connected[p.From()]
				if !ok {
					continue
				}
				if recipient, ok := s.connected[p.To()]; ok {
					recipient.Send() <- p
				} else {
					sender.Send() <- payload.NewRecipientNotConnected(p.From())
				}
			}
		}
	}
}
