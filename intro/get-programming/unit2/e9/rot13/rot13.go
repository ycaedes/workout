package main

import "fmt"

// Decipher the message
func main() {
	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c += 13
			if c > 'z' {
				c -= 26
			}
		}
		fmt.Printf("%c", c)
	}
}
