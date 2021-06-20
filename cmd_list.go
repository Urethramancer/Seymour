// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"strings"

	"github.com/Urethramancer/signor/opt"
)

// ListCmd options.
type ListCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"PODCAST" help:"Specifying the name of a podcast lists its episodes instead."`
}

func (cmd *ListCmd) Run(in []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if cmd.Name != "" {
		needle := strings.ToLower(cmd.Name)
		for k, pod := range list.List {
			haystack := strings.ToLower(k)
			if strings.Contains(haystack, needle) {
				fmt.Printf("%s - %d episodes:\n", pod.Name, len(pod.Episodes))
				for _, ep := range pod.Episodes {
					fmt.Printf("\t%s (%s)\n", ep.Name, ep.Updated)
				}
				println()
				return nil
			}
		}

		return unknownPodcast(cmd.Name)
	}

	for _, pod := range list.List {
		println(pod.Name)
	}
	return nil
}
