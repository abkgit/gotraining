package main

import "fmt"

func main() {

	fmt.Println(half(50))
	fmt.Println(half(25))
}

func half(num int) (float64, bool) {
	return float64(num) / 2, num%2 == 0
}
