package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/Urethramancer/Seymour/feed"
	"github.com/Urethramancer/cross"
	"github.com/Urethramancer/signor/stringer"
)

func fetchFeed(url string) (*feed.Feed, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	s := stringer.New()
	_, err = io.Copy(s, r.Body)
	if err != nil {
		return nil, err
	}

	rss, err := feed.NewRSS([]byte(s.String()))
	if err != nil {
		return nil, err
	}

	fn := feedFile(rss.Title)
	f, err := os.Create(fn)
	if err != nil {
		return nil, err
	}

	_, err = f.WriteString(s.String())
	if err != nil {
		return nil, err
	}

	return rss, nil
}

func podcastName(s string) string {
	fn := strings.ReplaceAll(s, " ", "")
	fn = strings.ReplaceAll(fn, "'", "")
	fn = strings.ReplaceAll(fn, "/", "")
	return fn
}

func podFile(s string) string {
	fn := podcastName(s)
	fn = filepath.Join(cross.ConfigPath(), podpath, fn)
	return fn
}

func feedFile(s string) string {
	fn := podcastName(s)
	fn = filepath.Join(cross.ConfigPath(), feedpath, fn)
	return fn
}
