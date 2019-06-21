package feed

import (
	"encoding/xml"
	"time"
)

// Feed represents a cached RSS/Atom feed.
type Feed struct {
	Title       string
	Subtitle    string
	Website     string
	Language    string
	Copyright   string
	Author      string
	Summary     string
	Description string
	Categories  string
	EpisodeList []*Episode
	Episodes    map[string]*Episode
	Date        time.Time
}

type RSSXML struct {
	XMLName  xml.Name  `xml:"rss"`
	Channels []Channel `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name   `xml:"channel"`
	Title       string     `xml:"title"`
	Website     string     `xml:"link"`
	Language    string     `xml:"language,omitempty"`
	Copyright   string     `xml:"copyright,omitempty"`
	Subtitle    string     `xml:"subtitle,omitempty"`
	Author      string     `xml:"author,omitempty"`
	Summary     string     `xml:"summary"`
	Description string     `xml:"description"`
	Explicit    string     `xml:"explicit"`
	Image       string     `xml:"image"`
	Categories  []string   `xml:"category"`
	Items       []*Episode `xml:"item"`
	Date        string     `xml:"pubDate,omitempty"`
}

type Episode struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	Author      string    `xml:"author,omitempty"`
	Subtitle    string    `xml:"subtitle,omitempty"`
	Summary     string    `xml:"summary"`
	Description string    `xml:"description"`
	URL         Enclosure `xml:"enclosure"`
	ActualDate  string    `xml:"pubDate"`
	Date        time.Time
}

type Enclosure struct {
	XMLName xml.Name
	URL     string `xml:"url,attr"`
	Length  int64  `xml:"length,attr"`
	Type    string `xml:"type,attr"`
}
