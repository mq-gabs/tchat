package main

import (
	"bufio"
	"fmt"
	"os"

	"tchat.com/client/command"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var cont = true
	var err error

	for cont {
		fmt.Printf("tchat > ")

		if !scanner.Scan() {
			fmt.Println("cannot scan input")
			return
		}

		input = scanner.Text()

		cont, err = command.Exec(input)
		if err != nil {
			fmt.Printf("ERROR: %v\n", err.Error())
		}
	}
}
