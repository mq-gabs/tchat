package cmd

import (
	"github.com/mq-gabs/kmdx"
	"tchat.com/client/config"
)

var (
	cmdExit   = "exit"
	cmdWhoAmI = "whoami"
	cmdEmpty  = ""

	cmdChat = "chat"
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
			whoAmI(conf.Me)
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
		c.Exec(func(s *kmdx.Scope) error {
			return startChat([]string{"blob"})
		})
	})

	return k
}
