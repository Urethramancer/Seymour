// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"sort"

	"github.com/Urethramancer/signor/opt"
)

// ListCmd options.
type ListCmd struct {
	opt.DefaultHelp
	Name string `placeholder:"PODCAST" help:"Lists the episodes of a podcast instead (or first partial match)."`
}

func (cmd *ListCmd) Run(in []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if cmd.Name != "" {
		pod := list.Find(cmd.Name)
		if pod == nil {
			return unknownPodcast(cmd.Name)
		}

		names := []string{}
		for _, ep := range pod.Episodes {
			names = append(names, fmt.Sprintf("%s (%s)", ep.Name, ep.Updated))
		}

		sort.Strings(names)
		fmt.Printf("%s - %d episodes:\n", pod.Name, len(pod.Episodes))
		for _, x := range names {
			fmt.Printf("\t%s\n", x)
		}

		println()
		return nil
	}

	names := []string{}
	count := 0
	for _, pod := range list.List {
		x := fmt.Sprintf("%s - %d episodes", pod.Name, len(pod.Episodes))
		count += len(pod.Episodes)
		names = append(names, x)
	}
	sort.Strings(names)
	for _, n := range names {
		println(n)
	}
	fmt.Printf("%d episodes total.\n", count)
	return nil
}
