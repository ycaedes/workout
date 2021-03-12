package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println("1.", terrestrial, gasGiants, iceGiants)
	fmt.Println("2.", gasGiants[0]) // index (window) into slice

	// take a slice of a slice
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]

	fmt.Println("3.", giants, gas, ice)

	// NOTE: modifying the element of a slice changes underlying
	// planets array.
	iceGiantsMarkII := iceGiants
	iceGiants[1] = "Poseidon"
	fmt.Println("4.", planets)
	fmt.Println("5.", iceGiants, iceGiantsMarkII, ice)

	// default indices for slicing
	terrestrial = planets[:4]
	iceGiants = planets[6:]
	allPlanets := planets[:] // slice containing all elements

	fmt.Println("6.", terrestrial, iceGiants, allPlanets)

	// slicing syntax for strings
	colonized := terrestrial[2:]
	fmt.Println("7.", colonized)
}
