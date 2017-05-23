package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x\n", 42, 42, 42)
	// fmt.Printf("%d - %b - %#x\n", 42, 42, 42) // %#x gives the 0x numeral notation
	// fmt.Printf("%d - %b - %#X\n", 42, 42, 42)   // capital %X for capital hex digits
	// fmt.Printf("%d \t %b \t %#X\n", 42, 42, 42) // tab escape char, \t
}
