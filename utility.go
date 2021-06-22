// Copyright (c) 2021 Ronny Bangsund
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package main

import (
	"bufio"
	"os"
	"strings"
)

func askString(q string) string {
	r := bufio.NewReader(os.Stdin)
	print(q)
	res, _ := r.ReadString('\n')
	return strings.ReplaceAll(res, "\n", "")
}
