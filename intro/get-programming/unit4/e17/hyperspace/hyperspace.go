package main

import (
	"fmt"
	"strings"
)

// hyperspace removes the space surrounding worlds
func hyperspace(worlds []string) {
	for i := range worlds {
		worlds[i] = strings.TrimSpace(worlds[i])
	}
}

// when passing arguments to functions, avoid using
// arrays directly. instead, pass the slice
func main() {
	planets := []string{" Venus		", "Earth	", " Mars"}
	hyperspace(planets)

	fmt.Println(strings.Join(planets, ""))
}

// NOTE: slices have a length, but unlike arrays,
// the length isn't apart of the type
