package chat

import (
	"fmt"
	"os/exec"

	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
)

type Chat struct {
	me      *users.User
	contact *users.User
	history string
}

func NewChat(me, contact *users.User) *Chat {
	return &Chat{
		me:      me,
		contact: contact,
		history: fmt.Sprintf("starting chat with %v\n", contact.Name),
	}
}

func (c *Chat) AddMessage(msg *messages.Message) {
	var name = "unknown"
	t := msg.SentAt.Format("2006-01-02 15:04:05")

	if msg.SentBy.ID == c.me.ID {
		name = "me"
	} else {
		name = string(msg.SentBy.Name)
	}

	c.history += fmt.Sprintf("%v (%v): %v\n", t, name, msg.Body)
}

func (c *Chat) Display() {
	bytes, _ := exec.Command("clear").Output()
	fmt.Println(string(bytes))
	fmt.Println(c.history)
}

func (c *Chat) LoadHistory(msgs []messages.Message) {
	for _, m := range msgs {
		c.AddMessage(&m)
	}
}
