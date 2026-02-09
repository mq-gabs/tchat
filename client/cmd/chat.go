package cmd

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"tchat.com/client/api"
	"tchat.com/client/chat"
	"tchat.com/client/cmd/cmdutils"
	"tchat.com/client/reader"
	"tchat.com/server/modules/users"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
)

func startChat(userID utils.UserID, chatApi *api.TChatAPI, sender *users.User) error {
	var err error

	receiver, err := chatApi.FindUserByID(&handlers.FindUserByIDQuery{
		UserID: userID,
	})
	if err != nil {
		return err
	}

	chatID, err := utils.MakeChatID(sender.ID, receiver.ID)
	if err != nil {
		return err
	}

	newMsgs, err := chatApi.WebsocketChat(chatID)
	if err != nil {
		return err
	}

	msgs, err := chatApi.ReadChat(&handlers.ReadChatQuery{
		User1: sender.ID,
		User2: receiver.ID,
	})
	if err != nil {
		return err
	}

	chat := chat.NewChat(sender, receiver)
	chat.LoadHistory(msgs)

	r := reader.New()

	cmdutils.EnterAlternateScreen()
	defer cmdutils.ExitAlternateScreen()

	wg := sync.WaitGroup{}
	chat.Display()
	fmt.Printf("> ")

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				return
			case nm := <-newMsgs:
				chat.AddMessage(nm)
				chat.Display()
				fmt.Printf("> ")
			}
		}
	}()

	go func() {
		defer cancel()
		defer wg.Done()

		cont := true

		for cont {
			input, err := r.Read()
			if err != nil {
				err = fmt.Errorf("cannot read input: %w", err)
				cont = false
				return
			}

			if strings.HasPrefix(input, "/exit") {
				cont = false
				err = nil
				break
			}

			if err = chatApi.SendMessage(&handlers.SendMessageBody{
				Body:       utils.MessageBody(input),
				SenderID:   sender.ID,
				ReceiverID: receiver.ID,
			}); err != nil {
				err = fmt.Errorf("cannot send message: %w", err)
				return
			}
		}
	}()

	wg.Add(2)
	wg.Wait()

	return err
}
