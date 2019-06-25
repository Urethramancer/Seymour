package main

import (
	"path/filepath"
	"time"

	"github.com/Urethramancer/cross"
)

// Config for Seymour.
type Config struct {
	// DownloadPath for the podcast episodes.
	DownloadPath string `json:"downloadpath"`
	// Frequency default.
	Frequency time.Duration `json:"frequency"`
}

func loadConfig() (*Config, error) {
	fn := filepath.Join(cross.ConfigPath(), "config.json")
	var cfg Config
	cfg.Frequency = time.Hour * 6
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
