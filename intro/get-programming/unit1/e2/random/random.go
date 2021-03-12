package main

import (
	"fmt"
	"math/rand"
)

// Generate pseudorandom numbers
func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	// Random distance between Earth and Mars
	num = rand.Intn(345000001) + 56000000
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)
}
