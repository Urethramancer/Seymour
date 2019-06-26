package feed

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
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

// NewRSSFromFile reads the entire feed file into memory and unmarshals it.
func NewRSSFromFile(name string) (*Feed, error) {
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}

	return NewRSS(data)
}

// NewRSS unmarshals a provided feed from XML.
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
		ep.ActualDate = fixDate(ep.ActualDate)
		t, err := time.Parse(time.RFC1123Z, ep.ActualDate)
		if err != nil {
			return nil, err
		}
		ep.Date = t
		f.EpisodeList = append(f.EpisodeList, ep)
		f.Episodes[ep.Title] = ep
	}
	var t time.Time
	if c.Date == "" {
		t = time.Now()
	} else {
		c.Date = fixDate(c.Date)
		t, err = time.Parse(time.RFC1123Z, c.Date)
	}
	if err != nil {
		return nil, err
	}
	f.Date = t
	return &f, err
}

func fixDate(s string) string {
	s = strings.Replace(s, "GMT", "+0000", 1)
	s = strings.Replace(s, "January", "Jan", 1)
	s = strings.Replace(s, "February", "Feb", 1)
	s = strings.Replace(s, "March", "Mar", 1)
	s = strings.Replace(s, "April", "Apr", 1)
	s = strings.Replace(s, "June", "Jun", 1)
	s = strings.Replace(s, "July", "Jul", 1)
	s = strings.Replace(s, "August", "Aug", 1)
	s = strings.Replace(s, "September", "Sep", 1)
	s = strings.Replace(s, "October", "Oct", 1)
	s = strings.Replace(s, "November", "Nov", 1)
	s = strings.Replace(s, "December", "Dec", 1)
	s = strings.Replace(s, " 1 ", " 01 ", 1)
	s = strings.Replace(s, " 2 ", " 02 ", 1)
	s = strings.Replace(s, " 3 ", " 03 ", 1)
	s = strings.Replace(s, " 4 ", " 04 ", 1)
	s = strings.Replace(s, " 5 ", " 05 ", 1)
	s = strings.Replace(s, " 6 ", " 06 ", 1)
	s = strings.Replace(s, " 7 ", " 07 ", 1)
	s = strings.Replace(s, " 8 ", " 08 ", 1)
	s = strings.Replace(s, " 9 ", " 09 ", 1)
	return s
}
