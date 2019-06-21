package main

import "github.com/Urethramancer/signor/log"

type CmdGetDownload struct{}

func (cmd *CmdGetDownload) Run(args []string) error {
	cfg, err := loadConfig()
	if err != nil || cfg.DownloadPath == "" {
		log.Default.Msg("No download path set.")
		return nil
	}

	log.Default.Msg("%s", cfg.DownloadPath)
	return nil
}
