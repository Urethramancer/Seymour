package main

import (
	"io"
	"net/http"

	"github.com/Urethramancer/Seymour/feed"
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
