package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

var Options struct {
	opt.DefaultHelp
	Podcast CmdAdd `command:"add" aliases:"a" help:"Add podcast."`
}

func main() {
	// l := log.Default

	a := opt.Parse(&Options)
	if Options.Help {
		a.Usage()
		return
	}
	a.RunCommand(false)

	// fi, err := os.Stat("smartenough")
	// if err != nil {
	// 	l.Err("Error loading file: %s", err.Error())
	// 	return
	// }

	// t := time.Now().Sub(fi.ModTime())
	// l.Msg("%f minutes old", t.Minutes())
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
