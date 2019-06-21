package feed

import (
	"encoding/xml"
	"errors"
	"reflect"
	"strings"
	"time"
)

func IsRSS(s string) bool {
	dec := xml.NewDecoder(strings.NewReader(s))
	for {
		t, err := dec.Token()
		if err != nil {
			return false
		}
		if reflect.TypeOf(t).String() == "xml.StartElement" {
			e := t.(xml.StartElement)
			if e.Name.Local == "rss" {
				return true
			}
		}
	}
}

func NewRSS(data []byte) (*Feed, error) {
	if !IsRSS(string(data)) {
		return nil, errors.New(("not an RSS feed"))
	}

	var x RSSXML
	err := xml.Unmarshal(data, &x)
	if err != nil {
		return nil, err
	}

	var f Feed
	c := x.Channels[0]
	f.Title = c.Title
	f.Subtitle = c.Subtitle
	f.Website = c.Website
	f.Language = c.Language
	f.Copyright = c.Language
	f.Author = c.Author
	f.Summary = c.Summary
	f.Description = c.Description
	f.Categories = strings.Join(c.Categories, ", ")
	f.Episodes = make(map[string]*Episode)
	for _, ep := range x.Channels[0].Items {
		t, err := time.Parse(time.RFC1123Z, ep.ActualDate)
		if err != nil {
			return nil, err
		}
		ep.Date = t
		f.EpisodeList = append(f.EpisodeList, ep)
		f.Episodes[ep.Title] = ep
	}
	t, err := time.Parse(time.RFC1123Z, c.Date)
	if err != nil {
		return nil, err
	}
	f.Date = t
	return &f, err
}

func NewRSSFromFile(filename string) (*Feed, error) {
	var feed Feed

	return &feed, nil
}
