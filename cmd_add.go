package main

import (
	"errors"
	"path/filepath"
	"strings"
	"time"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdAdd struct {
	opt.DefaultHelp
	URL string `placeholder:"URL" help:"URL to the feed."`
}

func (cmd *CmdAdd) Run(args []string) error {
	if cmd.Help || cmd.URL == "" {
		return errors.New(opt.ErrorUsage)
	}

	rss, err := fetchFeed(cmd.URL)
	if err != nil {
		return err
	}

	m := log.Default.Msg
	m("%s has %d episodes and was last updated %s", rss.Title, len(rss.EpisodeList), rss.Date.String())

	p := Podcast{
		Title:     rss.Title,
		URL:       cmd.URL,
		Updated:   rss.Date,
		Frequency: time.Hour * 6,
	}

	fn := strings.ReplaceAll(p.Title, " ", "-")
	fn = filepath.Join(cross.ConfigPath(), feedpath, fn)
	err = SaveJSON(fn, p)
	if err != nil {
		return err
	}

	m("Saved %s", fn)
	return nil
}
