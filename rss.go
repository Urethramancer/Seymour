// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"encoding/xml"
	"strings"
)

// RSS feed.
type RSS struct {
	XMLName  xml.Name  `xml:"rss"`
	Version  string    `xml:"version,attr"`
	Channels []Channel `xml:"channel"`
}

// Channel in an RSS feed.
type Channel struct {
	XMLName   xml.Name `xml:"channel"`
	Title     string   `xml:"title"`
	PubDate   string   `xml:"pubDate"`
	LastBuild string   `xml:"lastBuildDate"`
	Website   string   `xml:"link"`
	Episodes  []Item   `xml:"item"`
}

// Item (episode) in a channel.
type Item struct {
	XMLName     xml.Name  `xml:"item"`
	Title       string    `xml:"title"`
	PubDate     string    `xml:"pubDate"`
	Link        string    `xml:"link"`
	Description string    `xml:"description"`
	Enclosure   Enclosure `xml:"enclosure"`
	Episode     int       `xml:"episode"`
}

type Enclosure struct {
	XMLName xml.Name `xml:"enclosure"`
	Type    string   `xml:"type,attr"`
	URL     string   `xml:"url,attr"`
}

// FileName for the enclosure.
func (i Item) FileName() string {
	b := strings.Builder{}
	b.WriteString(i.Title)
	switch i.Enclosure.Type {
	case "audio/mpeg":
		b.WriteString(".mp3")
	}
	return b.String()
}
