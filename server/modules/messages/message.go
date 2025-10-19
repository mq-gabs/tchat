package messages

import (
	"time"

	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type Message struct {
	ID     utils.MessageID   `json:"id"`
	Body   utils.MessageBody `json:"body"`
	SentBy *users.User       `json:"sent_by"`
	SentTo *users.User       `json:"sent_to"`
	SentAt time.Time         `json:"sent_at"`
}

func New(body utils.MessageBody, sentBy, sentTo *users.User) *Message {
	return &Message{
		ID:     utils.MessageID(utils.NewID()),
		Body:   utils.MessageBody(body),
		SentBy: sentBy,
		SentTo: sentTo,
		SentAt: time.Now(),
	}
}
