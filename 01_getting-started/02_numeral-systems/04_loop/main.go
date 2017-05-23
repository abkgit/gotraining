package main

import "fmt"

// Prints hexadecimal and binary for decimal numbers upto 200
func main() {
	for i := 0; i <= 200; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
