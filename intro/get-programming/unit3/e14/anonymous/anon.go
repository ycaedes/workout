package main

import "fmt"

// declare anonymous functions in one step
func main() {
	func() {
		fmt.Println("Functions anonymous")
	}()
}
