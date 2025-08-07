package commands

import (
	"fmt"
	"errors"

	"github.com/curator4/gator/internal/config"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandMap map[string]func(*config.State, Command) error
}

func (c *Commands) Run(s *config.State, cmd Command) error {
	handler, ok := c.CommandMap[cmd.Name]
	if !ok {
		return fmt.Errorf("no command: '%s'", cmd.Name)
	}

	return handler(s, cmd)
}

func (c *Commands) Register(name string, f func(*config.State, Command) error) {
	c.CommandMap[name] = f
}

func HandlerLogin(s *config.State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return errors.New("must have a single argument: username")
	}

	if err := s.Cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("username set: %s\n", cmd.Args[0])
	return nil
}
