package main

import "fmt"

type kelvin float64
type celsius float64

// kelvinToCelsius converts °K to °C
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15) // type conversion
}

// celsiusToKelvin converts °C to °K
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := kelvinToCelsius(k)
	fmt.Println(k, "° K is ", c, "° C")

	c = 127.0
	k = celsiusToKelvin(c)
	fmt.Println(c, "° C is ", k, "° K")
}
