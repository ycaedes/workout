package main

import "fmt"

func main() {
	var planets [8]string

	planets[0] = "Mercury" // assign planet to index 0
	planets[1] = "Venus"
	planets[2] = "Earth"

	earth := planets[2] // retrieve planet at index 2
	fmt.Println(earth)
	fmt.Println(len(planets))
	fmt.Println(planets[3] == "")
}
