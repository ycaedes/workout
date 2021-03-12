package main

import "fmt"

// terraform func is variadic in the sense that it accepts a variable
// number of arguments denoted by ellipsis used in last parameter
func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds))

	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {
	twoWorlds := terraform("New", "Venus", "Mars")
	fmt.Println("1.", twoWorlds)

	// to pass a slice instead of multiple arguments, expand the slice
	// with an ellipsis
	planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println("2.", newPlanets)
}

// What are the three uses of the ellipsis in Go:
// 1. Have the Go compiler count the number of elements in a composite
//	  literal for an array.
// 2. Make the last parameter of a variadic function capture zero or more
//    arguments as a slice.
// 3. Expand the elements of a slice into arguments passed to a function.
