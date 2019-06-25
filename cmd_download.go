package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/Urethramancer/Seymour/feed"
	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/log"
	"github.com/Urethramancer/signor/opt"
)

type CmdDownload struct {
	opt.DefaultHelp
	// All episodes will be downloaded, no matter what.
	All bool `short:"a" long:"all" help:"Download all podcasts."`
	// Latest episode only will be downloaded.
	Latest bool `short:"l" long:"latest" help:"Download only the latest episode."`
	// Name of single podcast to download episodes from.
	Name string `placeholder:"NAME" help:"Podcast to download."`
}

func (cmd *CmdDownload) Run(args []string) error {
	if cmd.Help {
		return errors.New(opt.ErrorUsage)
	}

	m := log.Default.Msg
	cfg, err := loadConfig()
	if err != nil || cfg.DownloadPath == "" {
		m("Set the download path before downloading episodes.")
		return nil
	}

	if cmd.Name != "" {
		return downloadEpisodes(cmd.Name, cfg.DownloadPath, cmd.Latest, cmd.All)
	}

	fp := filepath.Join(cross.ConfigPath(), feedpath)
	files, err := ioutil.ReadDir(fp)
	if err != nil {
		return err
	}

	for _, f := range files {
		err = downloadEpisodes(f.Name(), cfg.DownloadPath, cmd.Latest, cmd.All)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadEpisodes(podcast, path string, latest, all bool) error {
	var p Podcast
	fn := feedFile(podcast)
	err := LoadJSON(fn, &p)
	if err != nil {
		return err
	}

	path = filepath.Join(path, podcast)
	if !cross.DirExists(path) {
		err = os.MkdirAll(path, 0700)
		if err != nil {
			return err
		}
	}

	m := log.Default.Msg
	m("Downloading %s (last updated %s)", p.Title, p.Updated.String())
	rss, err := fetchFeed(p.URL)
	if err != nil {
		return err
	}

	if latest {
		if rss.EpisodeList[0].Title == p.LastDownload {
			m("No new downloads since %s.", p.LastDownload)
			return nil
		}
		m("\tDownloading %s…", rss.EpisodeList[0].Title)
		err = downloadEpisode(rss.EpisodeList[0], path)
		if err != nil {
			return err
		}

		p.LastDownload = rss.EpisodeList[0].Title
		SaveJSON(fn, p)
		return nil
	}

	if all || p.LastDownload == "" {
		m("\tDownloading all %d episodes.", len(rss.EpisodeList))
		p.LastDownload = rss.EpisodeList[0].Title
		SaveJSON(fn, p)
		return nil
	}

	var dlist []*feed.Episode
	for _, ep := range rss.EpisodeList {
		if ep.Title == p.LastDownload {
			break
		}
		dlist = append(dlist, ep)
	}

	if len(dlist) == 0 {
		m("No new downloads since %s.", p.LastDownload)
		return nil
	}

	if len(dlist) == 1 {
		m("\tDownloading %s.", rss.EpisodeList[0].Title)
	} else {
		m("\tDownloading the %d newest episodes.", len(dlist))
	}

	p.LastDownload = dlist[0].Title
	m("Saving %s", fn)
	SaveJSON(fn, p)

	for _, ep := range dlist {
		err = downloadEpisode(ep, path)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadEpisode(ep *feed.Episode, path string) error {
	u, err := url.Parse(ep.URL.URL)
	if err != nil {
		return err
	}

	name := ep.Title + filepath.Ext(u.Path)
	fn := filepath.Join(path, name)

	r, err := http.Get(ep.URL.URL)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	size, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return err
	}

	src := &Counter{
		Reader: r.Body,
		Name:   name,
		Length: int64(size),
	}

	f, err := os.Create(fn)
	if err != nil {
		return err
	}

	io.Copy(f, src)
	fmt.Println("")
	return nil
}
