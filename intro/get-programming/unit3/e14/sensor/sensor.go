package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

// fakeSensor generates a pseudorandom number between
// 150-300° K
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// TODO: implement a real sensor
func realSensor() kelvin {
	return 0
}

func main() {
	sensor := fakeSensor // initialize function to var
	fmt.Println(sensor())

	sensor = realSensor // assign function to var
	fmt.Println(sensor())
}
