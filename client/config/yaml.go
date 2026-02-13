package config

import (
	"tchat.com/server/modules/users"
	"tchat.com/server/utils"
)

type ConfigOptions struct {
	Name    utils.UserName `yaml:"name"`
	Servers []ConfigServer `yaml:"servers,omitempty"`
}

type ConfigServer struct {
	Host    string       `yaml:"host"`
	User    users.User   `yaml:"user"`
	Friends []users.User `yaml:"friends"`
}
