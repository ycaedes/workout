package main

import (
	"fmt"
)

// Initializing structures with composite literals:
func main() {
	type location struct {
		lat, long float64
	}

	// opportunity and insight are initialized using field-value pairs.
	// fields may be in any order and fields that aren't listed will retain
	// the zero value for their type:
	opportunity := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	// you can prepend the %v format verb with a plus sign + to print out
	// the fields names, which is especially useful for inspecting large
	// structures:
	curiosity := location{-4.5895, 137.4417}
	fmt.Printf("%v\n", curiosity)
	fmt.Printf("%+v\n", curiosity)

	// structures are copied:
	bradbury := location{-4.5895, 137.4417}
	curiosity = bradbury

	curiosity.long += 0.0106
	fmt.Println(bradbury, curiosity)
}
