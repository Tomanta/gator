package main

import (
	//"fmt"
	"github.com/tomanta/gator/internal/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config: %v\n", err)
		return
	}

	st := &state{
		cfg: &cfg,
	}

	cmdList := commands{commandList: map[string]func(*state, command) error{}}
	cmdList.register("login", handlerLogin)

	args := os.Args

	if len(args) < 2 {
		log.Fatalf("No command provided")
	}

	cmd := command{
		name:      args[1],
		arguments: args[2:],
	}
	err = cmdList.run(st, cmd)
	if err != nil {
		log.Fatalf("Error running command!")
	}
}
