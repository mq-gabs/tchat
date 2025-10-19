package command

import (
	"errors"
	"os/exec"
)

func chat(args []string) (bool, error) {
	if len(args) == 0 {
		return true, errEmptyArgs
	}

	cmd := exec.Command("smcup")

	err := cmd.Run()

	if err != nil {
		return true, errors.Join(errCannotExec, err)
	}

	return true, nil
}
