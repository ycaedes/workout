package main

import "fmt"

// assign anon func to var
var f = func() {
	fmt.Println("Dress up for the masquerade.")
}

func main() {
	f()
}
