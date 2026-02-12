package main

import (
	"errors"
	"fmt"

	"tchat.com/client/api"
	"tchat.com/client/config"
	"tchat.com/client/reader"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
)

const (
	messageNotConfigured = `
welcome!
your tchat client is not configured...
please, type the server host
`
)

var (
	errUndefined = errors.New("undefined")
)

func setup() (*config.Config, error) {
	conf := &config.Config{}
	r := reader.New()

	if err := conf.Validate(); err == nil {
		return nil, errUndefined
	}

	fmt.Printf(messageNotConfigured)
	var host string

	for host == "" {
		fmt.Printf("\nhost: ")
		input, err := r.Read()
		if err != nil {
			return nil, err
		}

		testAPi := api.NewTChatAPI(input)
		if err := testAPi.Ping(); err != nil {
			fmt.Printf("invalid host: %v\n", err)
		} else {
			host = input
		}
	}

	chatApi := api.NewTChatAPI(host)

	fmt.Printf("\nnow type your username\n")

	var userName string
	for userName == "" {
		fmt.Printf("user name: ")
		input, err := r.Read()
		if err != nil {
			return nil, err
		}

		userName = input
	}

	u, err := chatApi.SaveUser(&handlers.SaveUserBody{
		Name: utils.UserName(userName),
	})
	if err != nil {
		return nil, err
	}

	newConf, err := config.New(u, chatApi)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\nlogged successfully!\nid: %v\nname: %v\n\n", u.ID, u.Name)

	return newConf, nil
}
