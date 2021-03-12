package main

import (
	"fmt"
	"sort"
)

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	// sorts planets alphabetically
	sort.StringSlice(planets).Sort()
	fmt.Println("1.", planets)

	// Strings helper func in sort package can perform
	// the type conversion and call Sort method
	sort.Strings(planets)
	fmt.Println("2.", planets)
}
