package service

import (
	"sync"

	"tchat.com/server/modules/messages"
	"tchat.com/server/utils"
)

type MessagesService struct {
	mu       sync.Mutex
	messages map[utils.ChatID]chan *messages.Message
}

func NewMessagesService() *MessagesService {
	return &MessagesService{
		messages: make(map[utils.ChatID]chan *messages.Message),
	}
}

func (ms *MessagesService) GetMessagesChannel(chatID utils.ChatID) chan *messages.Message {
	ch, ok := ms.messages[chatID]
	if ok {
		return ch
	}

	ms.mu.Lock()
	ms.messages[chatID] = make(chan *messages.Message)
	ms.mu.Unlock()

	return ms.messages[chatID]
}
