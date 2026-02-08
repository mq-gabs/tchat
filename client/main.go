package main

import (
	"errors"
	"fmt"

	"tchat.com/client/api"
	"tchat.com/client/cmd"
	"tchat.com/client/config"
	"tchat.com/client/reader"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
)

var (
	conf    *config.Config
	chatApi *api.TChatAPI
)

const (
	messageNotConfigured = `
welcome!
your tchat client is not configured...
please, type the server host
`
)

func init() {
	conf = config.New()
	r := reader.New()

	if err := conf.Validate(); err == nil {
		return
	}

	fmt.Printf(messageNotConfigured)
	var host string

	for host == "" {
		fmt.Printf("\nhost: ")
		input, err := r.Read()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		testAPi := api.NewTChatAPI(input)
		if err := testAPi.Ping(); err != nil {
			fmt.Printf("invalid host: %v\n", err)
		} else {
			host = input
		}
	}

	chatApi = api.NewTChatAPI(host)

	fmt.Printf("\nnow type your username\n")

	var userName string
	for userName == "" {
		fmt.Printf("user name: ")
		input, err := r.Read()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		userName = input
	}

	u, err := chatApi.SaveUser(&handlers.SaveUserBody{
		Name: utils.UserName(userName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	conf.Me = u

	fmt.Printf("\nlogged successfully!\nid: %v\nname: %v\n\n", u.ID, u.Name)
}

func main() {
	if conf == nil || chatApi == nil {
		return
	}

	r := reader.New()
	cli := cmd.Setup()

	for {
		fmt.Printf("tchat > ")

		input, err := r.Read()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		err = cli.Exec(input)
		if err == nil {
			continue
		}

		if errors.Is(err, cmd.ErrExit) {
			fmt.Println("Bye!")
			return
		}

		fmt.Printf("ERROR: %v\n", err.Error())

		if errors.Is(err, cmd.ErrFatal) {
			return
		}
	}
}
