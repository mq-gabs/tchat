package cmdutils

import (
	"errors"
	"fmt"
	"os/exec"
)

var (
	errCannotOpenOverScreen  = errors.New("cannot open overscreen")
	errCannotCloseOverScreen = errors.New("cannot close overscreen")
)

func EnterAlternateScreen() error {
	bytes, err := exec.Command("tput", "smcup").Output()
	if err != nil {
		return errors.Join(errCannotOpenOverScreen, err)
	}

	fmt.Println(string(bytes))

	return nil
}

func ExitAlternateScreen() error {
	bytes, err := exec.Command("tput", "rmcup").Output()
	if err != nil {
		return errors.Join(errCannotCloseOverScreen, err)
	}

	fmt.Println(string(bytes))

	return nil
}
