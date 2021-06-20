// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"fmt"
	"sort"
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
