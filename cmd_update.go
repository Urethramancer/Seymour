package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
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

	if cmd.Name != "" {
		return updatePodcast(cmd.Name)
	}

	fp := filepath.Join(cross.ConfigPath(), podpath)
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = updatePodcast(f.Name())
		if err != nil {
			return err
		}
	}
	return nil
}

func updatePodcast(name string) error {
	fn := podFile(name)
	var p Podcast
	err := LoadJSON(fn, &p)
	if err != nil {
		return err
	}

	m := log.Default.Msg
	m("%s:", p.Title)

	t := p.Updated.Add(p.Frequency)
	if !time.Now().After(t) {
		log.Default.Msg("\tUp to date.")
		return nil
	}

	rss, err := fetchFeed(p.URL)
	if err != nil {
		return err
	}

	p.Updated = time.Now()
	p.MostRecent = rss.EpisodeList[0].Title
	err = SaveJSON(fn, p)
	if err != nil {
		return err
	}

	m("\tUpdated.")
	return nil
}
