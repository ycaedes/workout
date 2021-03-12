package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Marshal func from json pkg is used to encode data in location
// into -> JSON format and returns JSON data as bytes:
func main() {
	type location struct {
		// use struct tags to specify the field names you want
		// json pkg to use:
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}
	curiosity := location{-4.5895, 137.4417}
	bytes, err := json.Marshal(curiosity)
	exitOnError(err)
	fmt.Println(string(bytes))
}

// exitOnError prints any errors and exits:
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// NOTE: from output, JSON keys match the field names of the location
// struct. For this to work, json pkg requires fields to be EXPORTED.
