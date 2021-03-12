package main

import "fmt"

// Decode UTF-8 encoded string with range keyword
func main() {
	question := "¿Cómo estás?"

	/*for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}*/

	// use blank identifier _, to ignore the index
	for _, c := range question {
		fmt.Printf("%c", c)
	}
}
