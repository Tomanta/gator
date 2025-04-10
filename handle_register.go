package main

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/tomanta/gator/internal/database"
	"os"
	"time"
)

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
