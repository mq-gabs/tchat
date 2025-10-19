package command

import "fmt"

func whoAmI() (bool, error) {
	fmt.Println("John Doe")

	return true, nil
}

func exit() (bool, error) {
	return false, nil
}

func empty() (bool, error) {
	return true, nil
}
