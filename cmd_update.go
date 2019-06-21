package main

import (
	"errors"

	"github.com/Urethramancer/signor/opt"
)

type CmdUpdate struct {
	opt.DefaultHelp
	Name string `placeholder:"URL" help:"Optional podcast name to update. All will be updated if unspecified."`
}

func (cmd *CmdUpdate) Run(args []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	return nil
}

func updatePodcast(name string) {

}
