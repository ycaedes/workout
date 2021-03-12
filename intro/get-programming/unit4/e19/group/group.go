package main

import (
	"fmt"
	"math"
)

// group temperatures together in divisions of 10° each:
func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// map with float64 keys and slices of float64's:
	groups := make(map[float64][]float64)

	for _, t := range temperatures {
		g := math.Trunc(t/10) * 10 // round temperatures down
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%v: %v\n", g, temperatures)
	}

	// makeshift set where each element is guaranteed to occur only once:
	set := make(map[float64]bool)
	for _, t := range temperatures {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	fmt.Println(set)
}
