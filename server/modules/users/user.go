package users

import "tchat.com/server/utils"

type User struct {
	ID   utils.UserID
	Name string
}

func New(name string) *User {
	return &User{
		ID:   utils.UserID(utils.NewID()),
		Name: name,
	}
}
