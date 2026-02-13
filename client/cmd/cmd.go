package cmd

import (
	"fmt"

	"github.com/mq-gabs/kmdx"
	"tchat.com/client/config"
	"tchat.com/server/utils"
)

var (
	cmdExit   = "exit"
	cmdWhoAmI = "whoami"
	cmdEmpty  = ""

	cmdChat = "chat"

	cmdServer        = "server"
	cmdServerAdd     = "add"
	cmdServerConnect = "connect"
)

func Setup(conf *config.Config) *kmdx.CLI {
	k := kmdx.New()

	k.Command(cmdExit, func(c *kmdx.Command) {
		c.Exec(func(s *kmdx.Scope) error {
			return exit()
		})
	})

	k.Command(cmdWhoAmI, func(c *kmdx.Command) {
		c.Exec(func(s *kmdx.Scope) error {
			whoAmI(conf.Me())
			return nil
		})
	})

	k.Command(cmdEmpty, func(c *kmdx.Command) {
		c.Exec(func(s *kmdx.Scope) error {
			empty()
			return nil
		})
	})

	k.Command(cmdChat, func(c *kmdx.Command) {
		var userID string

		c.Flags(func(fs kmdx.FlagSetter) {
			fs.String("userid", &userID)
		})

		c.Exec(func(s *kmdx.Scope) error {
			if !conf.IsConnected() {
				return errServerNotConnected
			}

			return startChat(utils.UserID(userID), conf.API(), conf.Me())
		})
	})

	k.Command(cmdServer, func(c *kmdx.Command) {
		c.Subcommand(cmdServerAdd, func(sc *kmdx.Command) {
			var serverHost string

			sc.Flags(func(fs kmdx.FlagSetter) {
				fs.String("host", &serverHost)
			})

			sc.Exec(func(s *kmdx.Scope) error {
				if serverHost == "" {
					return errServerHostMustBeProvided
				}

				if err := conf.AddServer(serverHost); err != nil {
					return err
				}

				fmt.Println("server created and already connected")

				return nil
			})
		})

		c.Subcommand(cmdServerConnect, func(sc *kmdx.Command) {
			var serverHost string

			sc.Flags(func(fs kmdx.FlagSetter) {
				fs.String("host", &serverHost)
			})

			sc.Exec(func(s *kmdx.Scope) error {
				if serverHost == "" {
					return errServerHostMustBeProvided
				}

				if err := conf.ConnectServer(serverHost); err != nil {
					return err
				}

				fmt.Println("server connected!")

				return nil
			})
		})
	})

	return k
}
