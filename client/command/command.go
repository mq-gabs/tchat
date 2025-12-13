package command

import "strings"

var (
	cmdExit   = "exit"
	cmdWhoAmI = "whoami"
	cmdEmpty  = ""

	cmdChat = "chat"
)

func Exec(input string) error {
	parts := strings.Split(input, " ")

	switch len(parts) {
	case 0:
		return errEmptyInput
	case 1:
		return execSingleCommand(parts[0])
	default:
		return execComplexCommand(parts[0], parts[1:])
	}
}

func execSingleCommand(command string) error {
	switch command {
	case cmdExit:
		return exit()
	case cmdWhoAmI:
		return whoAmI()
	case cmdEmpty:
		return empty()
	default:
		return errInvalidCommand
	}
}

func execComplexCommand(command string, args []string) error {
	switch command {
	case cmdChat:
		return startChat(args)
	default:
		return errInvalidCommand
	}
}
