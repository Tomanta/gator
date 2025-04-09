package main

import (
	"fmt"
)


func handlerLogin(s *state, cmd command) error {
	// If arg's slice is empty, return an error
	// Use state to access te config struct to set the user to the username, return any errors
	// Print a message to the terminal that the user has been set
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	newUser := cmd.arguments[0]

	err := s.cfg.SetUser(newUser)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set to %s\n", newUser)
	return nil
}
