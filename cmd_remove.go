package main

import (
	"errors"
	"os"

	"github.com/Urethramancer/cross"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

// CmdRemove deletes podcasts and feeds.
type CmdRemove struct {
	opt.DefaultHelp
	FeedOnly bool     `short:"f" long:"feedonly" help:"Only delete the downloaded feed."`
	Name     []string `placeholder:"PODCAST" help:"Name of podcast."`
}

// Run removal.
func (cmd *CmdRemove) Run(args []string) error {
	if cmd.Help || len(cmd.Name) == 0 {
		return errors.New(opt.ErrorUsage)
	}

	var err error
	m := log.Default.Msg
	for _, podcast := range cmd.Name {
		m("%s:", podcast)
		fn := feedFile(podcast)
		if cross.FileExists(fn) {
			err = os.Remove(fn)
			if err != nil {
				return err
			}
			m("Removed %s", fn)
		} else {
			m("Unknown podcast '%s'", podcast)
		}

		if cmd.FeedOnly {
			continue
		}

		fn = podFile(podcast)
		if cross.FileExists(fn) {
			err = os.Remove(fn)
			if err != nil {
				return err
			}
			m("Removed %s", fn)
		} else {
			m("Unknown podcast '%s'", podcast)
		}
	}
	m("\nDone.")
	return nil
}
