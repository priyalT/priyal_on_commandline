package main

import (
	"fmt"
	"os"

	"priyal_on_commandline/internal"
)

func main() {
	// 1. Path to your config file
	configPath := "config/config.json"

	// 2. Load the configuration
	cfg, err := internal.LoadConfig(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not load config from %s: %v\n", configPath, err)
		os.Exit(1)
	}

	// 3. Application Start
	fmt.Printf("🚀 Starting %s v%s\n", cfg.AppName, cfg.Version)
	fmt.Printf("Logged in as: %s\n", cfg.Author)
	fmt.Println("--------------------------------")
	
}