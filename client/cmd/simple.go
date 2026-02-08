package cmd

import (
	"fmt"

	"tchat.com/server/modules/users"
)

func whoAmI(u *users.User) error {
	fmt.Printf("\nid: %v\nname: %v\n\n", u.ID, u.Name)

	return nil
}

func exit() error {
	return ErrExit
}

func empty() error {
	return nil
}
