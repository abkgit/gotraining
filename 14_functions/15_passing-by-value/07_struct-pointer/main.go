package main

import "fmt"

type customer struct {
	name string
	age  int
}

func main() {
	c1 := customer{"Abhishek", 26}
	fmt.Println(&c1.name) // 0xc42000a380

	changeMe(&c1)

	fmt.Println(c1)       // {Todd 26}
	fmt.Println(&c1.name) // 0xc42000a380
}

func changeMe(z *customer) {
	fmt.Println(z)       // &{Abhishek 26}
	fmt.Println(&z.name) // 0xc42000a380
	z.name = "Todd"
	fmt.Println(z)       // &{Todd 26}
	fmt.Println(&z.name) // 0xc42000a380

}
