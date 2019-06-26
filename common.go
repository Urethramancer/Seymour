package main

import (
	"strconv"
	"strings"
	"text/scanner"
	"time"
	"unicode"

	"github.com/Urethramancer/signor/log"
)

// TimeSince a specific point.
type TimeSince struct {
	// Since when?
	Since string `short:"s" long:"since" help:"Date-time to list episodes since."`
}

// TimePeriod is last N <time units>.
type TimePeriod struct {
	// Period up to now.
	Period string `short:"p" long:"period" help:"Period of time to go back for episode lists."`
}

func parsePeriod(period string) time.Duration {
	var d time.Duration
	s := scanner.Scanner{}
	s.Init(strings.NewReader(period))
	var n int64
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		switch s.TokenText() {
		case "h", "hour", "hours":
			d += time.Hour * time.Duration(n)
		case "d", "day", "days":
			d += time.Hour * 24 * time.Duration(n)
		case "w", "week", "weeks":
			d += time.Hour * 24 * 7 * time.Duration(n)
		case "m", "month", "months":
			d += time.Hour * 24 * 30 * time.Duration(n)
		case "y", "year", "years":
			d += time.Hour * 24 * 365 * time.Duration(n)
		default:
			x, _ := strconv.Atoi(s.TokenText())
			n = int64(x)
		}
	}
	log.Default.Msg("%s", d.String())
	return d
}

func stripNonNumeric(s string) string {
	num := func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}
	return strings.Map(num, s)
}

func stripNonLetter(s string) string {
	num := func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return -1
	}
	return strings.Map(num, s)
}

func timeUnitWord(u string) string {
	switch u {
	case "s":
		return "seconds"
	case "m":
		return "minutes"
	case "h":
		return "hours"
	case "d":
		return "days"
	}
	return ""
}
