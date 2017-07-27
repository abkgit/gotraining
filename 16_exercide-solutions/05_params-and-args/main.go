package main

import "fmt"

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}

func foo(i ...int) {
	for _, v := range i {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
	//fmt.Println(i)	//Toddd's solution. Prints the slice received.
}
