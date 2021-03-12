package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2018
	switch month := rand.Intn(13); month {
	case 2:
		day := rand.Intn(29)
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(31)
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(32)
		fmt.Println(era, year, month, day)
	}
}
