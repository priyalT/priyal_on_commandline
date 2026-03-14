package internal

import (
	"encoding/json"
	"os"
)

// Config holds the application settings
type Config struct {
	AppName  string `json:"app_name"`
	Author   string `json:"author"`
	Version  string `json:"version"`
}

// LoadConfig reads a JSON file from the given path
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}