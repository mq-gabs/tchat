package config

import (
	"errors"
	"os"
	"sync"

	"tchat.com/client/api"
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

var configFileName = "tchat.yaml"

type Config struct {
	currentUserData *users.User
	api             *api.TChatAPI
	options         *ConfigOptions
	mu              sync.Mutex
}

func New() (*Config, error) {
	c := &Config{nil, nil, &ConfigOptions{}, sync.Mutex{}}

	if err := c.loadConfig(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Config) Me() *users.User {
	return c.currentUserData
}

func (c *Config) API() *api.TChatAPI {
	return c.api
}

func (c *Config) saveFile() {
	saveConf(c.options)
}

func (c *Config) IsSet() bool {
	return c.options.Name != ""
}

func (c *Config) UpdateName(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.options.Name = utils.UserName(name)
	c.saveFile()
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

		conf = &ConfigOptions{}
	}

	c.options = conf

	return nil
}
