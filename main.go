package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

const (
	program  = "Seymour"
	feedpath = "feeds"
)

var Options struct {
	opt.DefaultHelp
	Add      CmdAdd      `command:"add" aliases:"a" help:"Add podcast."`
	List     CmdList     `command:"list" aliases:"ls,l" help:"List podcasts."`
	Update   CmdUpdate   `command:"update" aliases:"up,u" help:"Update episode list for podcast(s)."`
	Download CmdDownload `command:"download" aliases:"dl" help:"Download episode(s) for podcast(s)."`
	Set      CmdSet      `command:"set" help:"Set configuration options."`
	Get      CmdGet      `command:"get" help:"Get configuration options."`
}

func init() {
	cross.SetConfigPath(program)
	fp := filepath.Join(cross.ConfigPath(), feedpath)
	if !cross.DirExists(fp) {
		err := os.MkdirAll(fp, 0755)
		if err != nil {
			log.Default.Err("Error creating %s: %s", cross.ConfigPath(), err.Error())
			os.Exit(2)
		}
	}
}

func main() {
	a := opt.Parse(&Options)
	if Options.Help {
		a.Usage()
		return
	}
	err := a.RunCommand(false)
	if err != nil {
		log.Default.Msg("Error: %s", err.Error())
	}
}

func download(url string) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	file, err := os.Create("output")
	if err != nil {
		return err
	}
	defer file.Close()

	n, err := io.Copy(file, r.Body)
	log.Default.Msg("Fetched %d bytes", n)
	return nil
}
