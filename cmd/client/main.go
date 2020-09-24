package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	whitespace = " \t\n"
)

const (
	helpText = `Commands:
  h - display this menu
  c - initiate chat with a user
  q - exit the application
`
)

func main() {
	// get user credentials
	var creds *creds
	for {
		c, err := getCreds(os.Stdin)
		if err != nil {
			log.Fatalln("failed to read credentials")
		}
		if c.username == "" || c.password == "" {
			fmt.Println("username or password cannot be empty")
			continue
		}
		creds = c
		break
	}
	// connect to backend websocket
	c, _ := NewTerminalClient("ws://localhost:5000", creds.username, creds.password)
	if err := c.Connect(); err != nil {
		log.Fatalln(err.Error())
	}
	defer c.Disconnect()
	// app loop
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		in, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal("error reading input")
		}
		in = strings.Trim(in, whitespace)
		switch in {
		case "c":
			fmt.Print("username: ")
			if in, err := reader.ReadString('\n'); err != nil {
				log.Fatal("error reading input")
			} else {
				c.Chat(strings.Trim(in, whitespace))
			}
		case "h":
			fmt.Println(helpText)
		case "q":
			fmt.Println("Bye!")
			goto loopExit
		default:
			fmt.Println("unrecognized command, try again")
			continue
		}
	}
loopExit:
}
