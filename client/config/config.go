package config

import (
	"tchat.com/client/api"
	"tchat.com/server/modules/users"
)

type Config struct {
	me  *users.User
	api *api.TChatAPI
}

func New(me *users.User, chatApi *api.TChatAPI) *Config {
	return &Config{me, chatApi}
}

func (c *Config) Validate() error {
	if c.me == nil {
		return errUserNotDefined
	}

	if c.api == nil {
		return errHostNotDefined
	}

	return nil
}

func (c *Config) Me() *users.User {
	return c.me
}

func (c *Config) API() *api.TChatAPI {
	return c.api
}
