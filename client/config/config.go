package config

import (
	"errors"
	"os"
	"sync"

	"tchat.com/client/api"
	"tchat.com/server/modules/users"
	"tchat.com/server/router/handlers"
	"tchat.com/server/utils"
)

var configFileName = "tchat.yaml"

type Config struct {
	currentUserData *users.User
	currentApi      *api.TChatAPI
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
	return c.currentApi
}

func (c *Config) saveFile() {
	saveConf(c.options)
}

func (c *Config) IsSet() bool {
	return c.options.Name != ""
}

func (c *Config) IsConnected() bool {
	return c.currentApi != nil
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

func (c *Config) GetServerByHost(host string) (ConfigServer, bool) {
	for _, s := range c.options.Servers {
		if s.Host == host {
			return s, true
		}
	}

	return ConfigServer{}, false
}

func (c *Config) AddServer(host string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.GetServerByHost(host); ok {
		return errHostAlreadyBeingUsed
	}

	chatAPI := api.NewTChatAPI(host)
	if err := chatAPI.Ping(); err != nil {
		return errors.Join(errCannotReachServerHost, err)
	}

	u, err := chatAPI.SaveUser(&handlers.SaveUserBody{
		Name: c.options.Name,
	})
	if err != nil {
		return errors.Join(errCannotSaveUser, err)
	}

	c.currentApi = chatAPI
	c.options.Servers = append(c.options.Servers, ConfigServer{
		Host: host,
		User: *u,
	})
	c.currentUserData = u

	c.saveFile()
	return nil
}

func (c *Config) ConnectServer(host string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	s, ok := c.GetServerByHost(host)
	if !ok {
		return errServerNotFound
	}

	chatApi := api.NewTChatAPI(s.Host)

	c.currentApi = chatApi
	c.currentUserData = &s.User

	c.saveFile()

	return nil
}
