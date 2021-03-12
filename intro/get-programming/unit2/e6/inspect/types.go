package main

import "fmt"

// Type inference
func main() {
	// [1] to use the first argument for the second
	// format verb instead of printing the variable twice
	days := 365.2425
	fmt.Printf("Type %T for %[1]v\n", days)

	a := "text"
	fmt.Printf("Type %T for %[1]v\n", a)

	b := 42
	fmt.Printf("Type %T for %[1]v\n", b)

	c := 3.14
	fmt.Printf("Type %T for %[1]v\n", c)

	d := true
	fmt.Printf("Type %T for %[1]v\n", d)
}
