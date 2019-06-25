package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/opt"
)

type CmdList struct {
	opt.DefaultHelp
	Full bool `short:"f" long:"full" help:"Full details listing."`
}

func (cmd *CmdList) Run(args []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	fp := filepath.Join(cross.ConfigPath(), feedpath)
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
