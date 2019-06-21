package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

type CmdSet struct {
	opt.DefaultHelp
	Download CmdSetDownload `command:"download" aliases:"dl" help:"Set download directory for podcast episodes."`
}

func (cmd *CmdSet) Run(args []string) error {
	return errors.New(opt.ErrorUsage)
}
