package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/tabwriter"
	"time"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdList struct {
	opt.DefaultHelp
	Full    bool   `short:"f" long:"full" help:"Full details listing."`
	Podcast string `placeholder:"PODCAST" help:"Optional podcast to list episodes from."`
	TimeSince
	TimePeriod
}

func (cmd *CmdList) Run(args []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	var err error
	if cmd.Podcast != "" {
		t := time.Time{}
		if cmd.Since != "" {
			t, err = time.Parse(time.RFC1123Z, cmd.Since)
			if err != nil {
				return err
			}
		}

		if cmd.Period != "" {
			d, err := time.ParseDuration(cmd.Period)
			if err != nil {
				return err
			}
			t = time.Now().Add(-d)
		}
		return listPodcast(cmd.Podcast, cmd.Full, t)
	}

	fp := filepath.Join(cross.ConfigPath(), podpath)
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		return err
	}

	tw := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)
	if cmd.Full {
		fmt.Fprintln(tw, "Filename:\tTitle:\tURL:\tUpdated:\tFrequency:")
	} else {
		fmt.Fprintln(tw, "Filename:\tTitle:\tUpdated:")
	}

	for _, x := range files {
		if x.IsDir() {
			continue
		}

		fn := filepath.Join(fp, x.Name())
		var p Podcast
		err := LoadJSON(fn, &p)
		if err != nil {
			return err
		}
		if cmd.Full {
			fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%s\n", x.Name(), p.Title, p.URL, p.Updated.Local().String(), p.Frequency.String())
		} else {
			fmt.Fprintf(tw, "%s\t%s\t%s\n", x.Name(), p.Title, p.Updated.Local().String())
		}
	}

	tw.Flush()
	return nil
}

func listPodcast(name string, full bool, since time.Time) error {
	log.Default.Msg("%s", since.String())
	return nil
}
