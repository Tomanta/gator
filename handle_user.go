package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/tomanta/gator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	// If arg's slice is empty, return an error
	// Use state to access te config struct to set the user to the username, return any errors
	// Print a message to the terminal that the user has been set
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	user, err := s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		fmt.Println("User does not exist!")
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User has been set to %s\n", user.Name)
	return nil
}

func handlerListUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error getting user list: %v\n", err)
	}

	currentUser := s.cfg.CurrentUserName

	for _, name := range users {
		if name == currentUser {
			fmt.Printf("* %s (current)\n", name)
		} else {
			fmt.Printf("* %s\n", name)
		}
	}
	return nil
}

func handlerReset(s *state, cmd command) error {

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error resetting: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Users table reset\n")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.arguments[0],
	}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		fmt.Println("User already exists!")
		os.Exit(1)
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Printf("User created! %v\n", user)
	return nil
}
