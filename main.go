package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tomanta/gator/internal/config"
	"github.com/tomanta/gator/internal/database"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading config: %v\n", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DB_URL)
	if err != nil {
		fmt.Println("Error connecting to database")
		os.Exit(1)
	}

	dbQueries := database.New(db)

	st := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmdList := commands{
		commandList: make(map[string]func(*state, command) error),
	}
	cmdList.register("login", handlerLogin)
	cmdList.register("register", handlerRegister)
	cmdList.register("reset", handlerReset)
	cmdList.register("users", handlerListUsers)

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
