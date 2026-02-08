package config

import "tchat.com/server/modules/users"

type Config struct {
	me   *users.User
	host string
}

func New(me *users.User, host string) *Config {
	return &Config{me, host}
}

func (c *Config) Validate() error {
	if c.me == nil {
		return errUserNotDefined
	}

	if c.host == "" {
		return errHostNotDefined
	}

	return nil
}

func (c *Config) Me() *users.User {
	return c.me
}
