package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/websocket"

	"github.com/navaz-alani/chat/pkg/payload"
)

type TerminalClient struct {
	username string
	password string
	baseUrl  string
	conn     *websocket.Conn
	in       chan payload.Payload
}

func NewTerminalClient(baseUrl, username, password string) (*TerminalClient, error) {
	return &TerminalClient{
		username: username,
		password: password,
		baseUrl:  baseUrl,
		in:       make(chan payload.Payload, 10),
	}, nil
}

func encode(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (c *TerminalClient) Connect() error {
	url := fmt.Sprintf("%s/new_ws", c.baseUrl)
	dialer := websocket.Dialer{
		Subprotocols:    []string{},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	header := http.Header{
		"Authorization": []string{"Basic " + encode(c.username, c.password)},
	}
	conn, _, err := dialer.Dial(url, header)
	if err != nil {
		return err
	}
	c.conn = conn
	go c.readIncoming()
	return nil
}

func (c *TerminalClient) Disconnect() error {
	return c.conn.Close()
}

func (c *TerminalClient) readIncoming() {
	for {
		p := new(payload.WsPayload)
		if err := c.conn.ReadJSON(p); err != nil {
			break
		}
		c.in <- p
	}
	close(c.in)
}

func display(p payload.Payload) {
	if p.IsError() {
		fmt.Printf("\tError > %s\n", p.Error().Error())
		return
	}
	fmt.Print("\tTHEM > ")
	fmt.Println(p.GetData())
}

func (c *TerminalClient) Chat(username string) error {
	fmt.Println("Press 'Ctrl + D' to end chat")
	fmt.Printf("Now chatting with '%s':\n", username)
	reader := bufio.NewReader(os.Stdin)
	for {
		// check if messages have been received
		if l := len(c.in); l > 0 {
			for i := 0; i < l; i++ {
				display(<-c.in)
			}
		}
		fmt.Print("\tYOU  > ")
		if in, err := reader.ReadString('\n'); err != nil {
			if err == io.EOF {
				fmt.Print("\n")
				fmt.Println("\tChat ended.")
				break
			} else {
				// abnormal read error
			}
		} else {
			in = strings.Trim(in, whitespace)
			if len(in) == 0 {
				continue
			}
			p := payload.NewTextMsg(username, c.username, in)
			if err := c.conn.WriteJSON(p); err != nil {
				return err
			}
		}
	}
	return nil
}
