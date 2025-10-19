package main

import (
	"fmt"

	"tchat.com/client/command"
)

func main() {
	var input string
	var cont = true
	var err error

	for cont {
		fmt.Printf("tchat > ")
		fmt.Scanln(&input)

		cont, err = command.Exec(input)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
		}
	}
}
