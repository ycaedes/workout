package main

import "fmt"

// Reusing structures with types
func main() {
	type location struct {
		lat  float64
		long float64
	}

	var spirit location
	spirit.lat = -14.5685
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity)
}
