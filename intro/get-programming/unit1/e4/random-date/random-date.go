package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func main() {
	year := 2019
	month := rand.Intn(13)
	daysInMonth := 31
	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}
	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}
