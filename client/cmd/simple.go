package cmd

import "fmt"

func whoAmI() error {
	fmt.Println("John Doe")

	return nil
}

func exit() error {
	return ErrExit
}

func empty() error {
	return nil
}
