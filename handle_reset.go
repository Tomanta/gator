package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		fmt.Printf("Error resetting: %w\n", err)
		os.Exit(1)
	}

	fmt.Printf("Users table reset\n")
	return nil
}
