package main

import "fmt"

func main() {

	a := 10
	b := "golang"
	c := 4.17
	d := true

	fmt.Printf("%T \n", a) // %T is a Go-syntax representation of the type of the value
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
}
