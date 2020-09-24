/*
wsClient.go defines a concrete websocket client for the chat application.
*/

package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/navaz-alani/chat/pkg/payload"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type wsClient struct {
	parentService ChatService
	username      string
	conn          *websocket.Conn
	incoming      chan payload.Payload
}

func NewWsClient(cs ChatService, w http.ResponseWriter, r *http.Request) (*wsClient, error) {
	// requestor provides username in basic auth format with upgrade request
	username, _, ok := r.BasicAuth()
	if !ok {
		return nil, fmt.Errorf("authentication not provided")
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &wsClient{
		parentService: cs,
		username:      username,
		conn:          conn,
		incoming:      make(chan payload.Payload),
	}, nil
}

func (c *wsClient) ID() string { return c.username }

func (c *wsClient) Send() chan<- payload.Payload {
	return c.incoming
}

func (c *wsClient) Listen() {
	go c.handleOutgoing()
	go c.handleIncoming()
}

func (c *wsClient) handleOutgoing() {
	// read from client and send to chatService for distribution
	for {
		p := new(payload.WsPayload)
		_, reader, err := c.conn.NextReader()
		if err != nil {
			// end both routines - error is permanent
			break
		}
		if err := json.NewDecoder(reader).Decode(p); err != nil {
			c.incoming <- payload.NewBadPayload(c.username)
		} else {
			c.parentService.Distribute() <- p
		}
	}
	// unregister participant and close operations
	c.parentService.RemoveParticipant(c)
	close(c.incoming)
}

func (c *wsClient) handleIncoming() {
	for {
		p, ok := <-c.incoming
		if !ok {
			// if channel has been closed, end routine
			break
		}
		// write to websocket connection
		c.conn.WriteJSON(p)
	}
}
