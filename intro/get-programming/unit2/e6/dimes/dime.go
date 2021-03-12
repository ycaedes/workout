package main

import "fmt"

// Add eleven dimes of type float64
func main() {
	piggyBank := 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	fmt.Println(piggyBank)
}
