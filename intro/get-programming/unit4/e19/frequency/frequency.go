package main

import "fmt"

func main() {
	temperatures := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// pre-allocated map using make function:
	frequency := make(map[float64]int)

	for _, t := range temperatures {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.2f occurs %d times\n", t, num)
	}
}

// NOTE: Pre-allocating a map with make can save the CPU some
// work later when the map get bigger.
