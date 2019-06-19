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
	EpisodeList []*Item
	Episodes    map[string]*Item
	Date        time.Time
}

type Episode struct {
	Title       string
	Subtitle    string
	Summary     string
	Description string
	URL         string
	Date        time.Time
}

type RSSXML struct {
	XMLName  xml.Name  `xml:"rss"`
	Channels []Channel `xml:"channel"`
}

type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Website     string   `xml:"link"`
	Language    string   `xml:"language"`
	Copyright   string   `xml:"copyright"`
	Subtitle    string   `xml:"subtitle"`
	Author      string   `xml:"author"`
	Summary     string   `xml:"summary"`
	Description string   `xml:"description"`
	Explicit    string   `xml:"explicit"`
	Image       string   `xml:"image"`
	Categories  []string `xml:"category"`
	Items       []Item   `xml:"item"`
	Date        string   `xml:"pubDate"`
}

type Item struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	Author      string    `xml:"author"`
	Subtitle    string    `xml:"subtitle"`
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
