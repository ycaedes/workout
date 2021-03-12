package main

import "fmt"

func main() {
	var pi rune = 960 // rune is an alias for int32 type
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33 // byte is an alias for uint8 type

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)
}
