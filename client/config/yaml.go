package config

import (
	"tchat.com/server/modules/users"
)

type ConfigOptions struct {
	Servers []ConfigServer `yaml:"servers"`
}

type ConfigServer struct {
	Host    string       `yaml:"host"`
	User    users.User   `yaml:"user"`
	Friends []users.User `yaml:"friends"`
}
