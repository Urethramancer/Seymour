package main

import (
	"fmt"
	"io"
)

// Counter wraps an io.Reader and prints progress.
type Counter struct {
	io.Reader
	// Name to display.
	Name string
	// Length to expect.
	Length int64
	count  int64
}

// Read wrapper to print progress.
func (c *Counter) Read(p []byte) (int, error) {
	n, err := c.Reader.Read(p)
	c.count += int64(n)
	fmt.Printf("\r\t%s: %d / %d bytes            ", c.Name, c.count, c.Length)
	return n, err
}
