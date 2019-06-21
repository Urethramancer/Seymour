package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdSetDownload struct {
	opt.DefaultHelp
	Path string `opt:"required" placeholder:"PATH" help:"Directory to store downloads in. A sub-directory will be created for each podcast."`
}

func (cmd *CmdSetDownload) Run(args []string) error {
	if cmd.Help || cmd.Path == "" {
		return errors.New(opt.ErrorUsage)
	}

	cfg, err := loadConfig()
	if err != nil {
		cfg = &Config{}
	}

	m := log.Default.Msg
	cfg.DownloadPath, err = filepath.Abs(cmd.Path)
	if err != nil {
		return err
	}

	if !cross.DirExists(cfg.DownloadPath) {
		err = os.MkdirAll(cfg.DownloadPath, 0755)
		if err != nil {
			return err
		}
	}

	err = cfg.save()
	if err != nil {
		return err
	}

	m("Download path set to %s", cfg.DownloadPath)
	return nil
}
