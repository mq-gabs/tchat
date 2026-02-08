package cmd

import (
	"fmt"
	"strings"
	"time"

	"tchat.com/client/api"
	"tchat.com/client/chat"
	"tchat.com/client/cmd/cmdutils"
	"tchat.com/client/reader"
	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
)

func startChat(userID utils.UserID, chatApi *api.TChatAPI, sender *users.User) error {

	receiver, err := chatApi.FindUserByID(&handlers.FindUserByIDQuery{
		UserID: userID,
	})
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

	var cont = true
	var input string
	r := reader.New()

	cmdutils.EnterAlternateScreen()
	defer cmdutils.ExitAlternateScreen()

	for cont {
		chat.Display()

		fmt.Printf("> ")

		input, err = r.Read()
		if err != nil {
			return err
		}

		if strings.HasPrefix(input, "/exit") {
			break
		}

		chat.AddMessage(&messages.Message{
			ID:     utils.MessageID(sender.ID),
			SentBy: sender,
			Body:   utils.MessageBody(input),
			SentTo: receiver,
			SentAt: time.Now(),
		})
	}

	return nil
}
