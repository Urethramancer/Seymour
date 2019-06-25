package main

import "github.com/Urethramancer/signor/log"

type CmdGetFrequency struct{}

func (cmd *CmdGetFrequency) Run(args []string) error {
	cfg, err := loadConfig()
	if err != nil || cfg.Frequency == 0 {
		log.Default.Msg("No update frequency set.")
		return nil
	}

	log.Default.Msg("%s", cfg.Frequency.String())
	return nil
}
