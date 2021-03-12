package main

import "fmt"

// Use constants to convert the distance from CMD
// galaxy to earth in light years
func main() {
	const distance = 236000000000000000
	const lightSpeed = 299793
	const secondsPerDay = 86400
	const daysPerYear = 365

	const years = distance / lightSpeed / secondsPerDay / daysPerYear

	fmt.Println("Canis Major Dwarf Galaxy is", years, "light years away.")
}
