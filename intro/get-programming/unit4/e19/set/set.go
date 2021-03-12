package main

import (
	"fmt"
	"sort"
)

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
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

	// map keys have arbitrary order, so they must be converted to a slice
	// before they can be sorted
	unique := make([]float64, 0, len(set))
	for t := range set {
		unique = append(unique, t)
	}
	sort.Float64s(unique)
	fmt.Println(unique)
}
