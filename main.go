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
		cfg: cfg,
	}

	commands := commands.Commands{
		var commandMap make(map[string]func(*config.State, commands.Command) error)
	}
	commands.Register("login", commands.HandlerLogin)

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "requies at least 2 args")
		os.Exit(1)
	}

	command := commands.Command{
		name: os.Args[0],
		args: os.Args[1:],
	}
	commands.Run(state, command)


	

	fmt.Printf("ConfigL %+v\n", cfg)
}
