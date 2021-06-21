// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grimdork/xos"
)

const cfgname = "podcasts.json"

func configName() string {
	cp, err := xos.NewConfig("Seymour")
	if err != nil {
		return cfgname
	}

	ok, _ := xos.Exists(cp.Path())
	if !ok {
		err = os.MkdirAll(cp.Path(), 0700)
		if err != nil {
			return cfgname
		}
	}
	fn := filepath.Join(cp.Path(), cfgname)
	return fn
}

func loadPodcastList() (*Podcasts, error) {
	data, err := os.ReadFile(configName())
	if err != nil {
		return nil, err
	}

	var list Podcasts
	err = json.Unmarshal(data, &list)
	if err != nil {
		return nil, err
	}

	return &list, nil
}

// Save tp JSON.
func (list *Podcasts) Save() error {
	data, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		return err
	}

	return os.WriteFile(configName(), data, 0600)
}

// SetDownloadPath configures the download directory and saves.
func (list *Podcasts) SetDownloadPath() error {
	r := bufio.NewReader(os.Stdin)
	if list.Path == "" {
		p, _ := os.UserHomeDir()
		list.Path = filepath.Join(p, "Podcasts")
	}

	fmt.Printf("Enter your preferred podcast download path [%s]: ", list.Path)
	newpath, _ := r.ReadString('\n')
	newpath = strings.ReplaceAll(newpath, "\n", "")
	if newpath != "" {
		list.Path = newpath
	}

	err := list.Save()
	if err != nil {
		return err
	}

	return os.MkdirAll(list.Path, 0700)
}
