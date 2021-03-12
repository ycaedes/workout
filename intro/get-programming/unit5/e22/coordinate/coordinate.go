package main

import (
	"fmt"
)

// coordinate in degrees, minutes, seconds:
type coordinate struct {
	d, m, s float64
	h       rune
}

// decimal method converts a DMS coordinate to decimal degrees:
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

// attaching methods to structures:
func main() {
	// Bradbury Landing DMS coordinate (4°35'22.2" S, 137°26'30.1" E):
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
