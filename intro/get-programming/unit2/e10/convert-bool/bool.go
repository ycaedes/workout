package main

import "fmt"

func main() {
	// Convert bool to string
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	// Convert string to bool
	yesNo = "no"
	launch = yesNo == "yes"
	fmt.Println("Ready for launch:", launch)

	// Convert bool to int
	launch = true
	var oneZero int
	if launch {
		oneZero = 1
	} else {
		oneZero = 0
	}
	fmt.Println("Ready for launch:", oneZero)
}
