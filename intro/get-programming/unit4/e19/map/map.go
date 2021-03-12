package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15, // composite literals are key-value pairs for maps
		"Mars":  -65,
	}
	temp := temperature["Earth"]
	fmt.Printf("1. On average the Earth is %v° C.\n", temp)
	temperature["Earth"] = 16
	temperature["Venus"] = 464
	fmt.Println("2.", temperature)

	// accessing a key that doesn't exist in the map results
	// to the zero value for the type(int):
	moon := temperature["Moon"]
	fmt.Println("3.", moon)

	// comma, ok: syntax to distinguish between non-existence of moon
	// and moon being present WITH temperature of 0° C:
	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("4. On average the moon is %v° C.\n", moon)
	} else {
		fmt.Println("5. Where is the moon?")
	}
}
