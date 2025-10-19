package store

import (
	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type Store interface {
	SaveUser(users.User) error
	FindUserByID(utils.UserID) (*users.User, error)
	SendMessage(body string, sentBy, sentTo *users.User) error
	ReadChat(user1, user2 *users.User) ([]*messages.Message, error)
}
