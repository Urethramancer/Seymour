package main

import "time"

// Podcast on-disk reference.
type Podcast struct {
	// Title of the podcast.
	Title string `json:"name"`
	// URL of the feed.
	URL string `json:"url"`
	// Updated time is fetched from the feed.
	Updated time.Time `json:"lastupdated,omitempty"`
	// Frequency of updates.
	Frequency time.Duration `json:"frequency"`
}
