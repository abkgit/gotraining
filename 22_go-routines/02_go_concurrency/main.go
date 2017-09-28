package main

import "fmt"

func main() {
	go foo()
	go bar()
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar:", i)
	}
}

// nothing happens
// func main starts a goroutine, then starts separate go routunes
// for foo() and bar() and exits concurrently. So, nothing is printed.
