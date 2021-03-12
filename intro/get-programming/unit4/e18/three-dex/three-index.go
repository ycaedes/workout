package main

import "fmt"

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")
	fmt.Println("1.", planets)

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println("2.", planets)

	_ = worlds
}
