package main

import (
	"fmt"
	"os"
	"github.com/tomanta/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	st := &state{
		cfg: &cfg,
	}

	cmdList := commands{
		commandList: make(map[string]func(*state, command) error),
	}
	cmdList.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("Usage: cli <command> [args...]\n")
		os.Exit(1)
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}
	err = cmdList.run(st, cmd)
	if err != nil {
		fmt.Printf("Error running command: %v\n", err)
		os.Exit(1)
	}
}
