package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"tchat.com/client/command"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var err error

	for {
		fmt.Printf("tchat > ")

		if !scanner.Scan() {
			fmt.Println("cannot scan input")
			return
		}

		input = scanner.Text()

		err = command.Exec(input)

		if err == nil {
			continue
		}

		if errors.Is(err, command.ErrExit) {
			fmt.Println("Bye!")
			return
		}

		fmt.Printf("ERROR: %v\n", err.Error())

		if errors.Is(err, command.ErrFatal) {
			return
		}
	}
}
