package main

import "fmt"

func main() {

	half := func(num int) (float64, bool) {
		return float64(num) / 2, num%2 == 0
	}

	fmt.Println(half(50))
	fmt.Println(half(25))
}
