package main

import "time"

type FeedEntry struct {
	Title   string    `json:"name"`
	URL     string    `json:"url"`
	Updated time.Time `json:"lastupdated,omitempty"`
}
