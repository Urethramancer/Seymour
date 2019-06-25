package main

import (
	"errors"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdAdd struct {
	opt.DefaultHelp
	URL []string `placeholder:"URL" help:"URL to the feed."`
}

func (cmd *CmdAdd) Run(args []string) error {
	if cmd.Help || len(cmd.URL) == 0 {
		return errors.New(opt.ErrorUsage)
	}

	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	for _, url := range cmd.URL {
		rss, err := fetchFeed(url)
		if err != nil {
			return err
		}

		m := log.Default.Msg
		m("%s has %d episodes and was last updated %s", rss.Title, len(rss.EpisodeList), rss.Date.String())

		p := Podcast{
			Title:      rss.Title,
			URL:        url,
			Updated:    rss.Date,
			Frequency:  cfg.Frequency,
			MostRecent: rss.EpisodeList[0].Title,
		}

		fn := feedFile(p.Title)
		err = SaveJSON(fn, p)
		if err != nil {
			return err
		}

		m("Saved %s", fn)
	}
	return nil
}
