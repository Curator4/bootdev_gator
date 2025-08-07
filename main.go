package main

import (
	"fmt"
	"os"

	"github.com/curator4/gator/internal/config"
	"github.com/curator4/gator/internal/commands"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to read config file", err)
		os.Exit(1)
	}
	state := config.State{
		Cfg: &cfg,
	}

	commands_struct := commands.Commands{
		CommandMap: make(map[string]func(*config.State, commands.Command) error),
	}
	commands_struct.Register("login", commands.HandlerLogin)

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "requies at least 2 args\n")
		os.Exit(1)
	}

	command := commands.Command{
		Name: os.Args[1],
		Args: os.Args[2:],
	}
	if err := commands_struct.Run(&state, command); err != nil {
		fmt.Fprintf(os.Stderr, "failed to run command %s : %v\n", command.Name, err)
		os.Exit(1)
	}
}
