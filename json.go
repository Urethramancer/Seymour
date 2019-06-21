package main

import (
	"encoding/json"
	"io/ioutil"
)

// LoadJSON loads a JSON file into a structure.
func LoadJSON(fn string, out interface{}) error {
	data, err := ioutil.ReadFile(fn)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, out)
}

// SaveJSON writes any structure to file as JSON.
func SaveJSON(fn string, data interface{}) error {
	var out []byte
	out, err := json.MarshalIndent(data, "\t", "")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fn, out, 0600)
}
