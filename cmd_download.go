// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"os"
	"path/filepath"
	"sort"

	"github.com/Urethramancer/signor/opt"
)

// DownloadCmd options.
type DownloadCmd struct {
	opt.DefaultHelp
	Podcast string `placeholder:"PODCAST" help:"Download eppisodes only from this podcast (or first partial match)."`
	Start   int    `short:"e" long:"episode" placeholder:"NUMBER" help:"Episode number to start at. Only applies when downloading a specific podcast."`
	Mark    bool   `short:"m" long:"mark" help:"Mark all skipped episodes as downloaded when starting on a specific one."`
	Force   bool   `short:"f" long:"force" help:"Force download of previously downloaded episodes."`
	Latest  bool   `short:"l" long:"latest" help:"Only download the latest episode."`
}

// Run the download command.
func (cmd *DownloadCmd) Run(in []string) error {
	if cmd.Help {
		return opt.ErrUsage
	}

	list := getPodcastList()
	if list.Path == "" {
		err := list.SetDownloadPath()
		if err != nil {
			return err
		}
	}

	if cmd.Podcast != "" {
		pod := list.Find(cmd.Podcast)
		if pod == nil {
			return unknownPodcast(cmd.Podcast)
		}

		path := filepath.Join(list.Path, pod.Name)
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}

		start := cmd.Start
		if cmd.Latest && len(pod.Episodes) > 0 {
			numbers := []int{}
			for _, e := range pod.Episodes {
				numbers = append(numbers, e.Number)
			}
			numbers = sort.IntSlice(numbers)
			start = numbers[len(numbers)-1]
		}
		if cmd.Mark {
			pod.MarkDownloaded(start)
		}
		pod.DownloadEpisodes(path, start, cmd.Force)
		return list.Save()
	}

	for _, pod := range list.List {
		path := filepath.Join(list.Path, pod.Name)
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
		start := cmd.Start
		if cmd.Latest && len(pod.Episodes) > 0 {
			numbers := []int{}
			for _, e := range pod.Episodes {
				numbers = append(numbers, e.Number)
			}
			numbers = sort.IntSlice(numbers)
			start = numbers[len(numbers)-1]
		}
		if cmd.Mark {
			pod.MarkDownloaded(start)
		}
		pod.DownloadEpisodes(path, start, cmd.Force)
	}

	return list.Save()
}
