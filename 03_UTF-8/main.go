package main

import "fmt"

// Prints hexadecimal, binary and UTF-8 chars for first 200 chars. (0 to 127 are old ASCII range)
func main() {
	for i := 0; i <= 200; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
		//%q = a single-quoted character literal safely escaped with Go syntax
	}
}
