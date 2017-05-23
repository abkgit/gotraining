package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
}

var y = 42 // order does not matter in package level scope being accessed in the inner block level scope
