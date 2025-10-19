package store

import (
	"fmt"

	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type TChatCache struct {
	users    map[utils.UserID]*users.User
	messages map[utils.Merged2UsersID][]*messages.Message
}

func NewCache() *TChatCache {
	return &TChatCache{
		users:    make(map[utils.UserID]*users.User),
		messages: make(map[utils.Merged2UsersID][]*messages.Message),
	}
}

func (c *TChatCache) SaveUser(u *users.User) error {
	c.users[u.ID] = u

	return nil
}
func (c *TChatCache) FindUserByID(id utils.UserID) (*users.User, error) {
	u, ok := c.users[id]
	if !ok {
		return nil, fmt.Errorf("%w: %v", errUserDoesNotExists, id)
	}

	return u, nil
}
func (c *TChatCache) SendMessage(body string, sentBy, sentTo *users.User) error {
	m := messages.New(body, sentBy, sentTo)
	mergedIds, err := utils.MergeIDs(string(sentBy.ID), string(sentTo.ID))
	if err != nil {
		return err
	}

	c.messages[mergedIds] = append(c.messages[mergedIds], m)

	return nil
}
func (c *TChatCache) ReadChat(user1, user2 *users.User) ([]*messages.Message, error) {
	mergedIds, err := utils.MergeIDs(string(user1.ID), string(user2.ID))
	if err != nil {
		return nil, err
	}

	ms := c.messages[mergedIds]

	return ms, nil
}
