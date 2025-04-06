package main

import (
	"fmt"

	"github.com/tomanta/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("Error reading: %v", err)
		return
	}

	cfg.SetUser("brian")

	cfg2, err := config.Read()

	if err != nil {
		fmt.Printf("Error reading 2nd time: %v", err)
		return
	}

	fmt.Printf("%v\n", cfg2)
}
