package main

import "fmt"

// Numeric type conversion via float64
func main() {
	age := 41
	marsAge := float64(age)

	marsDays := 687.0
	earthDays := 365.2425
	marsAge *= earthDays / marsDays
	fmt.Println("I am", marsAge, "years old on Mars.")
	fmt.Printf("I am %.2f years old on Mars.", marsAge)
}
