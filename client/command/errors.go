package command

import "errors"

var (
	errEmptyInput     = errors.New("empty input")
	errEmptyArgs      = errors.New("empty args")
	errInvalidCommand = errors.New("invalid command")
	errCannotExec     = errors.New("cannot exec command")
)
