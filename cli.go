package main

import (
	"errors"
	"fmt"

	"github.com/tomanta/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	commandList map[string]func(*state, command) error
}

func getCommands() map[string]func(*state, command) error {
	return map[string]func(*state, command) error{
		"login": handlerLogin,
	}
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandList[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	return c.commandList[cmd.name](s, cmd)
}

func handlerLogin(s *state, cmd command) error {
	// If arg's slice is empty, return an error
	// Use state to access te config struct to set the user to the username, return any errors
	// Print a message to the terminal that the user has been set
	if len(cmd.arguments) != 1 {
		return errors.New("Login command expects exactly one argument")
	}

	newUser := cmd.arguments[0]

	err := s.cfg.SetUser(newUser)
	if err != nil {
		return err
	}
	fmt.Printf("User has been set to %s", newUser)
	return nil
}
