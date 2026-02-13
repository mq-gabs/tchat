package main

import (
	"fmt"

	"tchat.com/client/config"
	"tchat.com/client/reader"
)

const (
	messageNotConfigured = `
welcome!
your tchat client is not configured...
please, type your user name
`
)

func setup() (*config.Config, error) {
	conf, err := config.New()
	if err != nil {
		return nil, err
	}

	if conf.IsSet() {
		return conf, nil
	}

	r := reader.New()

	fmt.Printf(messageNotConfigured)

	var userName string

	for userName == "" {
		fmt.Printf("user name: ")
		input, err := r.Read()
		if err != nil {
			return nil, err
		}

		userName = input
	}

	conf.UpdateName(userName)

	fmt.Printf("\nlet's start...\n\n")

	return conf, nil
}
