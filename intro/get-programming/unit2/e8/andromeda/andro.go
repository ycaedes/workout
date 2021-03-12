package main

import (
	"fmt"
	"math/big"
)

func main() {
	// NewInt takes int64 and returns big.Int
	lightSpeed := big.NewInt(299792)
	secondsPerDay := big.NewInt(86400)

	// SetString method to initialize value from string
	distance := new(big.Int)
	distance.SetString("24000000000000000000", 10)
	fmt.Println("Andromeda Galaxy is", distance, "km away.")

	seconds := new(big.Int)
	seconds.Div(distance, lightSpeed)

	days := new(big.Int)
	days.Div(seconds, secondsPerDay)

	fmt.Println("That is", days, "days of travel at light speed.")
}
