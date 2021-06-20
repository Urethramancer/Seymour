// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/xml"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

// Podcasts is a list of podcasts followed.
type Podcasts struct {
	// List is keyed to titles.
	List map[string]*Podcast `json:"podcasts"`
	Path string              `json:"downloadpath"`
}

type Podcast struct {
	// Name of podcast.
	Name string `json:"name"`
	// Website link to podcast.
	Website string `json:"website"`
	// RSS feed for podcast.
	RSS string `json:"rss"`
	// Episodes ordered from most recent to oldest.
	Episodes []Episode `json:"episodes"`
	// Downloaded episodes list.
	Downloaded map[int]bool
}

// Episode filename and download link.
type Episode struct {
	Name    string `json:"name"`
	Updated string `json:"updated"`
	Link    string `json:"link"`
	Number  int    `json:"number"`
}

func getPodcastList() *Podcasts {
	list, err := loadPodcastList()
	if err != nil {
		return &Podcasts{List: make(map[string]*Podcast)}
	}

	return list
}

// AddFeed from URL or file.
func (list *Podcasts) AddFeed(url string) (*Podcast, error) {
	data, err := WebDownload(url)
	if err != nil {
		return nil, err
	}

	var rss RSS
	d := xml.NewDecoder(strings.NewReader(string(data)))
	d.Strict = false
	err = d.Decode(&rss)

	if err != nil {
		return nil, err
	}

	// Most services only support one channel.
	c := rss.Channels[0]
	pod := &Podcast{
		Name:       c.Title,
		Website:    c.Website,
		RSS:        url,
		Downloaded: make(map[int]bool),
	}

	for _, ep := range rss.Channels[0].Episodes {
		e := Episode{
			Name:    fmt.Sprintf("%s %04d - %s", pod.Name, ep.Episode, ep.FileName()),
			Updated: ep.PubDate,
			Link:    ep.Enclosure.URL,
			Number:  ep.Episode,
		}
		pod.Episodes = append(pod.Episodes, e)
	}

	sort.SliceStable(pod.Episodes, func(i, j int) bool {
		return pod.Episodes[i].Number < pod.Episodes[j].Number
	})

	old, ok := list.List[pod.Name]
	if ok {
		pod.Downloaded = old.Downloaded
	}
	list.List[pod.Name] = pod
	return pod, nil
}

// DownloadEpisodes of podcast.
func (pod *Podcast) DownloadEpisodes(path string, start int, force bool) {
	for _, ep := range pod.Episodes {
		if ep.Number < start {
			continue
		}

		if pod.Downloaded[ep.Number] && !force {
			fmt.Printf("Skipping already downloaded episode %d\n", ep.Number)
			continue
		}

		fn := filepath.Join(path, ep.Name)
		fmt.Printf("Downloading %s:\n", fn)
		err := FileDownload(ep.Link, fn)
		println()
		if err != nil {
			fmt.Printf("Error downloading %s: %s\n", ep.Name, err.Error())
		} else {
			pod.Downloaded[ep.Number] = true
		}
	}
}

// MarkDownloaded marks all episodes up to this number as downloaded.
func (pod *Podcast) MarkDownloaded(n int) {
	for i := 1; i <= n; i++ {
		pod.Downloaded[i] = true
	}
}
