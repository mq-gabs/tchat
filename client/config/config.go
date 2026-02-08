package config

import "tchat.com/server/modules/users"

type Config struct {
	Me   *users.User
	Host string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Validate() error {
	if c.Me == nil {
		return errUserNotDefined
	}

	if c.Host == "" {
		return errHostNotDefined
	}

	return nil
}
