package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"tchat.com/client/chat"
	"tchat.com/client/cmd/cmdutils"
	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

func startChat(args []string) error {
	if len(args) == 0 {
		return errEmptyArgs
	}

	cmdutils.EnterAlternateScreen()

	sender := &users.User{
		ID:   "asldkfj",
		Name: "John",
	}
	receiver := &users.User{
		ID:   "laskdjfk",
		Name: "bob",
	}

	chat := chat.NewChat(sender, receiver)
	var cont = true
	var input string
	scanner := bufio.NewScanner(os.Stdin)

	for cont {
		chat.Display()

		fmt.Printf("> ")

		if !scanner.Scan() {
			return errCannotScanMessage
		}

		input = scanner.Text()

		if strings.HasPrefix(input, "/exit") {
			cmdutils.ExitAlternateScreen()
			break
		}

		chat.AddMessage(&messages.Message{
			ID:     "adjflsdakf",
			SentBy: sender,
			Body:   utils.MessageBody(input),
			SentTo: receiver,
			SentAt: time.Now(),
		})
	}

	return nil
}
