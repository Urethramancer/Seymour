// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"

	"github.com/Urethramancer/signor/opt"
)

// UpdateCmd options.
type UpdateCmd struct {
	opt.DefaultHelp
	Name    string `placeholder:"PODCAST" help:"Specifying the name of a podcast updates only that."`
	Replace bool   `short:"r" long:"replace" help:"Replace the episode list with the feed contents (remove links to episodes no longer listed)."`
}

func (cmd *UpdateCmd) Run(in []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if cmd.Name != "" {
		pod := list.Find(cmd.Name)
		if pod == nil {
			return unknownPodcast(cmd.Name)
		}

		if cmd.Replace {
			_, err := list.AddFeed(pod.RSS)
			if err != nil {
				return err
			}
		} else {
			err := pod.Update()
			if err != nil {
				return err
			}
		}

		fmt.Printf("Updated %s\n", pod.Name)
		return list.Save()
	}

	for _, pod := range list.List {
		if cmd.Replace {
			_, err := list.AddFeed(pod.RSS)
			if err != nil {
				return err
			}
		} else {
			err := pod.Update()
			if err != nil {
				return err
			}
		}

		fmt.Printf("Updated %s\n", pod.Name)
	}

	return list.Save()
}
