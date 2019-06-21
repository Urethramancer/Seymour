package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

type CmdGet struct {
	opt.DefaultHelp
	Download CmdGetDownload `command:"download" aliases:"dl" help:"Set download directory for podcast episodes."`
}

func (cmd *CmdGet) Run(args []string) error {
	return errors.New(opt.ErrorUsage)
}
