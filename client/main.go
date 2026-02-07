package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"tchat.com/client/cmd"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var input string
	var err error

	cli := cmd.Setup()

	for {
		fmt.Printf("tchat > ")

		if !scanner.Scan() {
			fmt.Println("FATAL: cannot scan input")
			return
		}

		input = scanner.Text()

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
