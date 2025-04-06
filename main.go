package main

import (
	"fmt"
	"log"

	"github.com/tomanta/gator/internal/config"
)

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
