package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

type CmdSet struct {
	opt.DefaultHelp
	Download CmdSetDownload  `command:"download" aliases:"dl" help:"Set download directory for podcast episodes."`
	Time     CmdSetFrequency `command:"frequency" aliases:"freq,fr" help:"Set frequency of feed updates."`
}

func (cmd *CmdSet) Run(args []string) error {
	return errors.New(opt.ErrorUsage)
}
