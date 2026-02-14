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
	cmdServerList    = "list"
	cmdServerConnect = "connect"

	cmdFriend     = "friend"
	cmdFriendAdd  = "add"
	cmdFriendList = "list"
	cmdFriendChat = "chat"
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

		c.Subcommand(cmdServerList, func(sc *kmdx.Command) {
			sc.Exec(func(s *kmdx.Scope) error {
				conf.ListServers()
				return nil
			})
		})

		c.Subcommand(cmdServerConnect, func(sc *kmdx.Command) {
			var serverIndex int

			sc.Flags(func(fs kmdx.FlagSetter) {
				fs.Int("i", &serverIndex)
			})

			sc.Exec(func(s *kmdx.Scope) error {
				if err := conf.ConnectServer(serverIndex); err != nil {
					return err
				}

				fmt.Println("server connected!")

				return nil
			})
		})
	})

	k.Command(cmdFriend, func(c *kmdx.Command) {
		c.Subcommand(cmdFriendAdd, func(sc *kmdx.Command) {
			var userID string

			sc.Flags(func(fs kmdx.FlagSetter) {
				fs.String("userid", &userID)
			})

			sc.Exec(func(s *kmdx.Scope) error {
				if userID == "" {
					return fmt.Errorf("%w: userid", errFlagNotDefined)
				}

				return conf.AddFriend(utils.UserID(userID))
			})
		})

		c.Subcommand(cmdFriendList, func(sc *kmdx.Command) {
			sc.Exec(func(s *kmdx.Scope) error {
				return conf.ListFriends()
			})
		})

		c.Subcommand(cmdFriendChat, func(sc *kmdx.Command) {
			var friendIndex int

			sc.Flags(func(fs kmdx.FlagSetter) {
				fs.Int("i", &friendIndex)
			})

			sc.Exec(func(s *kmdx.Scope) error {
				f, err := conf.GetFriendByIndex(friendIndex)
				if err != nil {
					return err
				}

				return startChat(f.ID, conf.API(), conf.Me())
			})
		})
	})

	return k
}
