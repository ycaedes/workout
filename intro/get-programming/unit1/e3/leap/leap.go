package main

import "fmt"

// Determine whether 2100 will be a leap year

// Leap Year:
// Any year evenly divisible by 4 but not evenly
// divisible by 100 or any year that is evenly divisible by 400
func main() {
	fmt.Println("The year is 2100, should you leap?")

	var year = 2000
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}
