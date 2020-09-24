package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

type creds struct {
	username string
	password string
}

func getCreds(r io.Reader) (*creds, error) {
	reader := bufio.NewReader(r)
	fmt.Println("Enter credentials now")
	fmt.Print("username: ")
	username, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	fmt.Print("password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, err
	}
	fmt.Print("\n\n")
	return &creds{
		username: strings.Trim(username, " \t\n"),
		password: string(bytePassword),
	}, nil
}
