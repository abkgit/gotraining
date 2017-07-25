package main

import "fmt"

func main() {

	name := "Abhishek"
	fmt.Println(&name) // 0xc42000e270

	changeMe(&name)

	fmt.Println(&name) // 0xc42000e270
	fmt.Println(name)  // Todd
}

func changeMe(z *string) {
	fmt.Println(z)  // 0xc42000e270
	fmt.Println(*z) // Abhishek
	*z = "Todd"
	fmt.Println(z)  // 0xc42000e270
	fmt.Println(*z) // Todd
}
