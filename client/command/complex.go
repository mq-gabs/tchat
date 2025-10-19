package command

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"tchat.com/client/chat"
	"tchat.com/server/modules/messages"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

func startChat(args []string) (bool, error) {
	if len(args) == 0 {
		return true, errEmptyArgs
	}

	bytes, err := exec.Command("tput", "smcup").Output()
	if err != nil {
		return true, errors.Join(errCannotExec, err)
	}
	fmt.Println(string(bytes))

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
			return true, errors.New("cannot scan message")
		}

		input = scanner.Text()

		if strings.HasPrefix(input, "/exit") {
			bytes, _ := exec.Command("tput", "rmcup").Output()
			fmt.Println(string(bytes))
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

	return true, nil
}
