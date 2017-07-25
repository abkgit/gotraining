package main

import "fmt"

func main() {

	name := "Abhishek"
	fmt.Println(name) // Abhishek

	changeMe(name)

	fmt.Println(name) // Abhishek
}

func changeMe(z string) {
	fmt.Println(z) // Abhishek
	z = "Todd"
	fmt.Println(z) // Todd
}
