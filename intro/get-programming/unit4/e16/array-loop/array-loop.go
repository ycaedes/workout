package main

import "fmt"

// Iterating through array via for/range
func main() {
	dwarfs := [5]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	// use traditional for loop when iterating in reverse
	// or to provide access to a specific element
	for i := 0; i < len(dwarfs); i++ {
		dwarf := dwarfs[i]
		fmt.Println(i, dwarf)
	}

	fmt.Println()

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

	// use blank identifier (underscore) if you don't need
	// the index var provided by range keyword
	for _, planet := range planets {
		fmt.Println(planet)
	}
}
