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
	Name string `placeholder:"PODCAST" help:"Specifying the name of a podcast updates only that."`
}

func (cmd *UpdateCmd) Run(in []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if cmd.Name != "" {
		pod, ok := list.List[cmd.Name]
		if !ok {
			return unknownPodcast(cmd.Name)
		}

		pod, err := list.AddFeed(pod.RSS)
		if err != nil {
			return err
		}

		fmt.Printf("Updated %s\n", pod.Name)
		return list.Save()
	}

	for _, pod := range list.List {
		pod, err := list.AddFeed(pod.RSS)
		if err != nil {
			return err
		}

		fmt.Printf("Updated %s\n", pod.Name)
	}

	return list.Save()
}
