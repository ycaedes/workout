package main

import "fmt"

type celsius float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

type fahrenheit float64

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)
}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

// function type declaration
type getRowFn func(row int) (string, string)

// drawTable displays a two column table
func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cellOne, cellTwo := getRow(row)
		fmt.Printf(rowFormat, cellOne, cellTwo)
	}
	fmt.Println(line)
}

func celsiusToFahrenheit(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenheit()
	cellOne := fmt.Sprintf(numberFormat, c)
	cellTwo := fmt.Sprintf(numberFormat, f)
	return cellOne, cellTwo
}

func fahrenheitToCelsius(row int) (string, string) {
	f := fahrenheit(row*5 - 40)
	c := f.celsius()
	cellOne := fmt.Sprintf(numberFormat, f)
	cellTwo := fmt.Sprintf(numberFormat, c)
	return cellOne, cellTwo
}

func main() {
	drawTable("°C", "°F", 29, celsiusToFahrenheit)
	fmt.Println()
	drawTable("°F", "°C", 29, fahrenheitToCelsius)
}
