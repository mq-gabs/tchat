package store

import (
	"fmt"
	"sync"

	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type TChatCache struct {
	users    map[utils.UserID]*users.User
	messages map[utils.MergedIDs][]*messages.Message
	mu       *sync.Mutex
}

func NewCache() *TChatCache {
	return &TChatCache{
		users:    make(map[utils.UserID]*users.User),
		messages: make(map[utils.MergedIDs][]*messages.Message),
	}
}

func (c *TChatCache) SaveUser(u *users.User) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.users[u.ID] = u

	return nil
}
func (c *TChatCache) FindUserByID(id utils.UserID) (*users.User, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	u, ok := c.users[id]
	if !ok {
		return nil, fmt.Errorf("%w: %v", ErrUserDoesNotExists, id)
	}

	return u, nil
}
func (c *TChatCache) SendMessage(m *messages.Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	mergedIds, err := utils.MergeIDs(string(m.SentBy.ID), string(m.SentTo.ID))
	if err != nil {
		return err
	}

	c.messages[mergedIds] = append(c.messages[mergedIds], m)

	return nil
}
func (c *TChatCache) ReadChat(user1, user2 *users.User) ([]*messages.Message, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	mergedIds, err := utils.MergeIDs(string(user1.ID), string(user2.ID))
	if err != nil {
		return nil, err
	}

	ms := c.messages[mergedIds]

	return ms, nil
}
