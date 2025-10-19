package messages

import (
	"time"

	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type Message struct {
	ID     utils.MessageID
	Body   string
	SentBy *users.User
	SentTo *users.User
	SentAt time.Time
}

func New(body string, sentBy, sentTo *users.User) *Message {
	return &Message{
		ID:     utils.MessageID(utils.NewID()),
		Body:   body,
		SentBy: sentBy,
		SentTo: sentTo,
		SentAt: time.Now(),
	}
}
