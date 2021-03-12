package main

import "fmt"

// Determine how fast a ship travels in order to reach
// Malacandra (Mars) in 28 days
func main() {
	const hoursPerDay = 24

	var days = 28
	var distance = 56000000 // km

	fmt.Println(distance/(days*hoursPerDay), "km/h")
}
