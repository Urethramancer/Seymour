// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Urethramancer/signor/opt"
)

// AddCmd options.
type AddCmd struct {
	opt.DefaultHelp
	URL  string `placeholder:"URL" help:"Address of feed to add. If a podcast with the same title exists, the feed is replaced."`
	List bool   `short:"l" long:"list" help:"The URL is instead a file with a URL per line to import feeds from."`
}

func (cmd *AddCmd) Run(in []string) error {
	if cmd.Help || cmd.URL == "" {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if cmd.List {
		data, err := os.ReadFile(cmd.URL)
		if err != nil {
			return err
		}

		urls := strings.Split(string(data), "\n")
		for _, u := range urls {
			if u == "" {
				continue
			}

			fmt.Printf("Adding %s\n", u)
			pod, err := list.AddFeed(u)
			if err != nil {
				return err
			}

			fmt.Printf("Added %s with %d episodes.\n\n", pod.Name, len(pod.Episodes))
		}
	} else {
		pod, err := list.AddFeed(cmd.URL)
		if err != nil {
			return err
		}

		fmt.Printf("Added %s with %d episodes.\n", pod.Name, len(pod.Episodes))
	}

	err := list.Save()
	if err != nil {
		return err
	}

	return nil
}
