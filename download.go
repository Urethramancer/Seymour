package main

import (
	"io"
	"net/http"
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

	return feed.NewRSS([]byte(s.String()))
}

func podcastName(s string) string {
	fn := strings.ReplaceAll(s, " ", "-")
	fn = strings.ReplaceAll(s, "'", "-")
	return fn
}

func feedFile(s string) string {
	fn := podcastName(s)
	fn = filepath.Join(cross.ConfigPath(), feedpath, fn)
	return fn
}
