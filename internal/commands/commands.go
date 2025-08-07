package commands

import (
	"fmt"
	"errors"

	"github.com/curator4/gator/config"
)

type Command struct {
	name string
	args []string
}

type Commands struct {
	commandMap map[string]func(*config.State, command) error
}

func (c *commands) Run(s *config.State, cmd command) error {
	handler, ok := c.commandMap[cmd.name]
	if !ok {
		return fmt.Errorf("no command: '%s'", cmd.name)
	}

	return handler(s, cmd)
}

func (c *commands) Register(name string, f func(*config.State, command) error) {
	c.commandMap[name] = f
	return nil
}

func HandlerLogin(s *config.State, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("must have a single argument: username")
	}

	if err := s.config.SetUser(cmd.args[0]); err != nil {
		return err
	}

	fmt.Printf("username set: %s\n", cmd.args[0])
	return nil
}
