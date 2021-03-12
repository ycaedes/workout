package main

import "fmt"

// Print each byte of "shalom" one char per line
func main() {
	message := "shalom"
	for i := 0; i < 6; i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}
}
