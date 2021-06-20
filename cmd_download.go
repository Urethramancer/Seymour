// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Urethramancer/signor/opt"
)

// DownloadCmd options.
type DownloadCmd struct {
	opt.DefaultHelp
	Podcast string `placeholder:"PODCAST" help:"Download eppisodes only from this podcast. Partial match allowed."`
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
		needle := strings.ToLower(cmd.Podcast)
		for k, pod := range list.List {
			haystack := strings.ToLower(k)
			if strings.Contains(haystack, needle) {
				path := filepath.Join(list.Path, pod.Name)
				err := os.MkdirAll(path, 0700)
				if err != nil {
					return err
				}
				pod.DownloadEpisodes(path, cmd.Start, cmd.Force)
				return list.Save()
			}
		}

		return unknownPodcast(cmd.Podcast)
	}

	for _, pod := range list.List {
		path := filepath.Join(list.Path, pod.Name)
		err := os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
		latest := 0
		if cmd.Latest && len(pod.Episodes) > 0 {
			latest = pod.Episodes[len(pod.Episodes)-1].Number
		}
		pod.DownloadEpisodes(path, latest, cmd.Force)
	}

	return list.Save()
}
