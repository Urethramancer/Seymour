// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import "fmt"

func unknownPodcast(name string) error {
	return fmt.Errorf("unknown podcast %s", name)
}
