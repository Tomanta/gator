package main

import (
	"fmt"
	"log"

	"github.com/tomanta/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

func (c *commands) register(name string, f func(*state, command) error) {

}

func (c *commands) run(name string, f func(*state, command) error) {

}

func handlerLogin(s *state, cmd command) error {
	// If arg's slice is empty, return an error
	// Use state to access te config struct to set the user to the username, return any errors
	// Print a message to the terminal that the user has been set
	return nil
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
		return
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("brian")

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Error reading 2nd time: %v\n", err)
		return
	}

	fmt.Printf("Read config again: %+v\n", cfg)
}
