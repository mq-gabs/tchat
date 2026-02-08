package main

import (
	"errors"
	"fmt"

	"tchat.com/client/cmd"
	"tchat.com/client/reader"
)

func main() {
	conf, err := setup()
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	r := reader.New()
	cli := cmd.Setup(conf)

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
