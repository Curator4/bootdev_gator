package config

import (
	"os"
	"fmt"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DatabaseURL string `json:"db_url"`
	UserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	configPath, err := getConfigPath()
	if err != nil {
		return Config{}, err
	}

	byteContent, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(byteContent, &cfg); err != nil {
		return Config{}, fmt.Errorf("unmarshaling config json %w", err)
	}

	return cfg, nil
}

func (cfg *Config) SetUser(name string) error {
	cfg.UserName = name
	if err := write(*cfg); err != nil {
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

func write(cfg Config) error {
	byteContent, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("marshalling config json: %w", err)
	}

	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(configPath, byteContent, 0644); err != nil {
		return fmt.Errorf("writing config file: %w", err)
	}

	return nil
}

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting home directory: %w", err)
	}

	return filepath.Join(homeDir, configFileName), nil
}
