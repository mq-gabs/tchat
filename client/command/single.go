package command

import "fmt"

func whoAmI() (bool, error) {
	fmt.Println("John Doe")

	return true, nil
}
