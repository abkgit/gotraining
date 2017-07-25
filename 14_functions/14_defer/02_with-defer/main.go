package main

import "fmt"

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // runs just before when the containing function, main, is about to return
	hello()
}
