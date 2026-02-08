package cmd

import "errors"

var (
	ErrExit              = errors.New("exit")
	ErrFatal             = errors.New("fatal")
	errEmptyInput        = errors.New("empty input")
	errEmptyArgs         = errors.New("empty args")
	errInvalidCommand    = errors.New("invalid command")
	errCannotExec        = errors.New("cannot exec command")
	errCannotScanMessage = errors.New("cannot scan message")
)
