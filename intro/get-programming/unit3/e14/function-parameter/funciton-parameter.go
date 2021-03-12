package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin // declare function type

// measureTemperature periodically calls fakeSensor func
// to log temperature data every second
func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

// fakeSensor generates a pseudorandom temperature of type
// kelvin between 150-300° K
func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func main() {
	measureTemperature(3, fakeSensor)
}
