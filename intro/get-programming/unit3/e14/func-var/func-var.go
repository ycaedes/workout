package main

import "fmt"

// declared var can be in scope of package or
// within a function
func main() {
	f := func(message string) {
		fmt.Println(message)
	}
	f("Go to the party.")
}
