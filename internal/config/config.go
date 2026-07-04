package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Trigger struct {
		Mode string `toml:"mode"`
	} `toml:"trigger"`
	Hotkey struct {
		Key int `toml:"key"`
	} `toml:"hotkey"`
	Heuristic struct {
		Value int `toml:"value"`
	} `toml:"heuristic"`
	Formatters struct {
		Order []string `toml:"order"`
	} `toml:"formatters"`
}

const configFilePath = ".config/pastiche/pastiche.toml"

func getConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, configFilePath), nil
}

func Load() (*Config, error) {
	confPath, err := getConfigPath()
	if err != nil {
		return nil, fmt.Errorf("getting config path: %w", err)
	}

	if _, err := os.Stat(confPath); errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	data, err := os.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("unmarshalling toml: %w", err)
	}

	return &config, nil
}
