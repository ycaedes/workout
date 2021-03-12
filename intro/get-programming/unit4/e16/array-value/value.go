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

	planetsMarkII := planets // COPIES planets array
	planets[2] = "Dripland"

	fmt.Println(planets) // prints modified array
	fmt.Println(planetsMarkII)
}
