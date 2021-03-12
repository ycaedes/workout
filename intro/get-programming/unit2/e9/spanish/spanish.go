package main

import (
	"fmt"
	"unicode/utf8"
)

// Determine length of string in runes and
// decode first char of string
func main() {
	question := "¿Cómo estás?"

	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")

	c, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune: %c %v bytes", c, size)
}
