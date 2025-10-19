package command

import "errors"

var (
	errEmptyInput     = errors.New("empty input")
	errInvalidCommand = errors.New("invalid command")
)
