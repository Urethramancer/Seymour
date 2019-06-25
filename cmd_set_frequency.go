package main

import (
	"errors"
	"time"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdSetFrequency struct {
	opt.DefaultHelp
	Time string `opt:"required" placeholder:"TIME" help:"Minimum time period before updating podcast feeds."`
}

func (cmd *CmdSetFrequency) Run(args []string) error {
	if cmd.Help || cmd.Time == "" {
		return errors.New(opt.ErrorUsage)
	}

	log.Default.Msg("Update every %v", cmd.Time)

	cfg, err := loadConfig()
	if err != nil {
		cfg = &Config{}
	}

	cfg.Frequency, err = time.ParseDuration(cmd.Time)
	if err != nil {
		return err
	}

	err = cfg.save()
	if err != nil {
		return err
	}

	log.Default.Msg("Default update frequency set to %s.", cfg.Frequency.String())
	return nil
}
