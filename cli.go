package main

import (
	"errors"
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
	f, ok := c.commandList[cmd.name]

	if !ok {
		return errors.New("command not found")
	}
	return f(s, cmd)
}
