package main

import "fmt"

func main() {
	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	// planets and planetsMarkII share the same underlying data,
	// so beware of the fact that changes to one impact the other:
	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets)
	fmt.Println(planetsMarkII)

	// when delete removes an element from the map, both planets
	// and planetsMarkII are impacted by the change:
	delete(planets, "Earth")
	fmt.Println(planetsMarkII)
}
