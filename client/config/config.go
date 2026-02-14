package config

import (
	"errors"
	"fmt"
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

func (c *Config) getCurrentServer() (*ConfigServer, error) {
	s, ok := c.GetServerByHost(c.currentApi.Host())
	if !ok {
		return nil, errServerNotFound
	}

	return s, nil
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

func (c *Config) GetServerByHost(host string) (*ConfigServer, bool) {
	for _, s := range c.options.Servers {
		if s.Host == host {
			return s, true
		}
	}

	return nil, false
}

func (c *Config) GetServerById(index int) (*ConfigServer, error) {
	if len(c.options.Servers) <= index {
		return nil, errInvalidIndex
	}

	return c.options.Servers[index], nil
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
	c.options.Servers = append(c.options.Servers, &ConfigServer{
		Host: host,
		User: *u,
	})
	c.currentUserData = u

	c.saveFile()
	return nil
}

func (c *Config) ConnectServer(index int) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	s, err := c.GetServerById(index)
	if err != nil {
		return errServerNotFound
	}

	chatApi := api.NewTChatAPI(s.Host)

	c.currentApi = chatApi
	c.currentUserData = &s.User

	c.saveFile()

	return nil
}

func (c *Config) ListServers() {
	if len(c.options.Servers) == 0 {
		fmt.Println("no servers registered")
		return
	}

	for i, s := range c.options.Servers {
		fmt.Printf("#%v: %v\n", i, s.Host)
	}
}

func (c *Config) GetFriendByServer(id utils.UserID, server *ConfigServer) (*users.User, bool) {
	for _, u := range server.Friends {
		if u.ID == id {
			return u, true
		}
	}

	return nil, false
}

func (c *Config) AddFriend(userID utils.UserID) error {
	s, ok := c.GetServerByHost(c.currentApi.Host())
	if !ok {
		return errServerNotFound
	}

	_, ok = c.GetFriendByServer(userID, s)
	if ok {
		return errFriendAlreadyAdded
	}

	u, err := c.currentApi.FindUserByID(&handlers.FindUserByIDQuery{
		UserID: userID,
	})
	if err != nil {
		return errors.Join(errUserNotFound, err)
	}

	s.Friends = append(s.Friends, u)

	c.saveFile()

	return nil
}

func (c *Config) GetFriendByIndex(index int) (*users.User, error) {
	s, err := c.getCurrentServer()
	if err != nil {
		return nil, err
	}

	if len(s.Friends) <= index {
		return nil, errInvalidIndex
	}

	return s.Friends[index], nil
}

func (c *Config) ListFriends() error {
	if !c.IsConnected() {
		return errNoServerIsConnected
	}

	s, ok := c.GetServerByHost(c.currentApi.Host())
	if !ok {
		return errServerNotFound
	}

	if len(s.Friends) == 0 {
		fmt.Println("friends list is empty")
		return nil
	}

	for i, f := range s.Friends {
		fmt.Printf("#%v:\n\tID: %v\n\tName: %v\n", i, f.ID, f.Name)
	}

	return nil
}
