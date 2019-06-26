package main

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
