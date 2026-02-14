package cmd

import "errors"

var (
	ErrExit                     = errors.New("exit")
	ErrFatal                    = errors.New("fatal")
	errEmptyInput               = errors.New("empty input")
	errEmptyArgs                = errors.New("empty args")
	errInvalidCommand           = errors.New("invalid command")
	errCannotExec               = errors.New("cannot exec command")
	errCannotScanMessage        = errors.New("cannot scan message")
	errServerHostMustBeProvided = errors.New("server host must be provided")
	errServerNotConnected       = errors.New("server not connected")
	errFlagNotDefined           = errors.New("flag not defined")
)
