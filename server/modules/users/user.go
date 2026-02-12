package users

import "tchat.com/server/utils"

type User struct {
	ID   utils.UserID   `json:"id" yaml:"id"`
	Name utils.UserName `json:"name" yaml:"name"`
}

func New(name utils.UserName) *User {
	return &User{
		ID:   utils.UserID(utils.NewID()),
		Name: utils.UserName(name),
	}
}
