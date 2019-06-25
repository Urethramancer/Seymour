package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

type CmdGet struct {
	opt.DefaultHelp
	Download CmdGetDownload  `command:"download" aliases:"dl" help:"Get download directory for podcast episodes."`
	Time     CmdGetFrequency `command:"frequency" aliases:"freq,fr" help:"Get frequency of feed updates."`
}

func (cmd *CmdGet) Run(args []string) error {
	return errors.New(opt.ErrorUsage)
}
