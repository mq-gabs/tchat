package config

import "errors"

var (
	errInvalidIndex = errors.New("invalid index")

	errUserNotDefined    = errors.New("user not defined")
	errHostNotDefined    = errors.New("host not defined")
	errOptionsNotDefined = errors.New("options not defined")

	errCannotReadConfigFile      = errors.New("cannot read config file")
	errCannotCreateConfigFile    = errors.New("cannot create config file")
	errCannotUnmarshalConfigFile = errors.New("cannot unmarshal config file")

	errHostAlreadyBeingUsed  = errors.New("host already being used")
	errCannotReachServerHost = errors.New("cannot reach server host")
	errCannotSaveUser        = errors.New("cannot save user")
	errServerNotFound        = errors.New("server not found")

	errFriendAlreadyAdded = errors.New("fried already added")
	errUserNotFound       = errors.New("user not found")
)
