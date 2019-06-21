package main

import (
	"path/filepath"

	"github.com/Urethramancer/cross"
)

// Config for Seymour.
type Config struct {
	DownloadPath string `json:"downloadpath"`
}

func loadConfig() (*Config, error) {
	fn := filepath.Join(cross.ConfigPath(), "config.json")
	var cfg Config
	err := LoadJSON(fn, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) save() error {
	fn := filepath.Join(cross.ConfigPath(), "config.json")
	return SaveJSON(fn, cfg)
}
