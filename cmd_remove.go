// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"

	"github.com/Urethramancer/signor/opt"
)

// RemoveCmd options.
type RemoveCmd struct {
	opt.DefaultHelp
	Podcast string `placeholder:"PODCAST" help:"Name of podcast to remove (exact match)."`
}

func (cmd *RemoveCmd) Run(in []string) error {
	if cmd.Help || cmd.Podcast == "" {
		return opt.ErrUsage
	}

	list := getPodcastList()
	pod := list.Find(cmd.Podcast)
	if pod == nil {
		return unknownPodcast(cmd.Podcast)
	}

	a := askString(fmt.Sprintf("Really remove %s [y/N]? ", pod.Name))
	if a != "y" && a != "Y" {
		return nil
	}

	fmt.Printf("Removed %s\n", pod.Name)
	delete(list.List, pod.Name)
	return list.Save()
}
