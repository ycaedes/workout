package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// display the JSON encoding of the three rover landing sites:
func main() {
	// slice of structures []struct is a collection of zero or more
	// values where each value is based on a structure instead of a primitive
	// type like float64:
	type location struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	// MarshallIndent applies indent to format the output:
	bytes, err := json.MarshalIndent(locations, "", " ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(bytes))
}
