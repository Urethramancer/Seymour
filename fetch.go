// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// WebDownload fetches the body of a URL.
func WebDownload(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New(res.Status)
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}

	return body, nil
}

// FileDownload fetches a file a chunk at a time and saves it.
func FileDownload(url, filename string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	defer res.Body.Close()
	f, err := os.Create(filename)
	if err != nil {
		return err
	}

	defer f.Close()
	c := Counter{Max: res.Header.Get("Content-Length")}
	_, err = io.Copy(f, io.TeeReader(res.Body, &c))
	if err != nil {
		return err
	}

	return nil
}

// Counter is a download progress filter.
type Counter struct {
	Total uint64
	Max   string
}

func (c *Counter) Write(p []byte) (int, error) {
	n := len(p)
	c.Total += uint64(n)
	fmt.Printf("\rDownloaded %d/%s bytes", c.Total, c.Max)
	return n, nil
}
