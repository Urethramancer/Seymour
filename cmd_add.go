package main

import (
	"errors"
	"io"
	"net/http"

	"github.com/Urethramancer/Seymour/feed"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
	"github.com/Urethramancer/signor/stringer"
)

type CmdAdd struct {
	opt.DefaultHelp
	URL string `placeholder:"URL" help:"URL to the feed."`
}

func (cmd *CmdAdd) Run(args []string) error {
	if cmd.Help || cmd.URL == "" {
		return errors.New(opt.ErrorUsage)
	}

	rss, err := downloadPodcast(cmd.URL)
	if err != nil {
		return err
	}
	m := log.Default.Msg
	m("%s has %d episodes and was last updated %s", rss.Title, len(rss.EpisodeList), rss.Date.String())

	return nil
}

func downloadPodcast(url string) (*feed.Feed, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	s := stringer.New()
	_, err = io.Copy(s, r.Body)
	if err != nil {
		return nil, err
	}
	f, err := feed.NewRSS([]byte(s.String()))
	return f, err
}
