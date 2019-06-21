package main

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdDownload struct {
	opt.DefaultHelp
	All    bool   `short:"a" long:"all" help:"Download all podcasts."`
	Latest bool   `short:"l" long:"latest" help:"Download only the latest episode."`
	Name   string `placeholder:"NAME" help:"Podcast to download."`
}

func (cmd *CmdDownload) Run(args []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	m := log.Default.Msg
	cfg, err := loadConfig()
	if err != nil || cfg.DownloadPath == "" {
		m("Set the download path before downloading episodes.")
		return nil
	}

	fp := filepath.Join(cross.ConfigPath(), feedpath)
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		return err
	}

	for _, f := range files {
		m("%s", f.Name())
	}
	return nil
}
