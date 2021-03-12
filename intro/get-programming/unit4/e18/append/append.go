package main

import (
	"fmt"
)

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)

	// append func is variadic, so you can pass multiple elements
	// at a time
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}
