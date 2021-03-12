package main

import (
	"fmt"
	"math"
)

// Demonstrate floating point precision 64 vs 32 bit
func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi
	fmt.Println(pi64)
	fmt.Println(pi32)
}
