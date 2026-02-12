package config

import (
	"errors"
	"os"

	"tchat.com/client/api"
	"tchat.com/server/modules/users"
)

var configFileName = "tchat.yaml"

type Config struct {
	me      *users.User
	api     *api.TChatAPI
	options *ConfigOptions
}

func New(me *users.User, chatApi *api.TChatAPI) (*Config, error) {

	c := &Config{me, chatApi, &ConfigOptions{}}

	if err := c.loadConfig(); err != nil {
		return nil, err
	}

	return c, nil
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

func (c *Config) loadConfig() error {
	conf, err := loadConfig()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	if err != nil {
		if err := createFile(); err != nil {
			return err
		}
	}

	c.options = conf

	return nil
}
