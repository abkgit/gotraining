package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string)	// Trying to use assertion without an interface
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string\n")
	}
}
