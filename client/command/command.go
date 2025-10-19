package command

import "strings"

var (
	cmdExit   = "exit"
	cmdWhoAmI = "whoami"
	cmdEmpty  = ""

	cmdChat = "chat"
)

func Exec(input string) (bool, error) {
	parts := strings.Split(input, " ")

	switch len(parts) {
	case 0:
		return true, errEmptyInput
	case 1:
		return execSingleCommand(parts[0])
	default:
		return execComplexCommand(parts[0], parts[1:])
	}
}

func execSingleCommand(command string) (bool, error) {
	switch command {
	case cmdExit:
		return exit()
	case cmdWhoAmI:
		return whoAmI()
	case cmdEmpty:
		return empty()
	default:
		return true, errInvalidCommand
	}
}

func execComplexCommand(command string, args []string) (bool, error) {
	switch command {
	case cmdChat:
		return chat(args)
	default:
		return true, errInvalidCommand
	}
}
