package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150) // type conversion
}

// calibrate accepts sensor as param and returns
// a replacement function
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 2
	sensor := calibrate(fakeSensor, offset)

	for count := 0; count < 10; count++ {
		fmt.Println(sensor())
	}
}
