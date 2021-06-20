package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/signor/opt"
)

var o struct {
	opt.DefaultHelp
	Add      AddCmd      `command:"add" help:"Add a podcast."`
	Remove   RemoveCmd   `command:"remove" aliases:"rm" help:"Remove a podcast."`
	List     ListCmd     `command:"list" aliases:"ls" help:"List podcasts or episodes."`
	Update   UpdateCmd   `command:"update" aliases:"up" help:"Update podcast(s)."`
	Download DownloadCmd `command:"downlaod" aliases:"dl" help:"Download podcast episodes."`
}

func main() {
	a := opt.Parse(&o)
	if o.Help {
		a.Usage()
		return
	}

	err := a.RunCommand(false)
	if err != nil {
		if err == opt.ErrNoCommand {
			a.Usage()
			return
		}

		fmt.Printf("Error running command: %s\n", err.Error())
		os.Exit(2)
	}
}
