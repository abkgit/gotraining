package main

import "fmt"

func main() {
	var expr bool = (true && false) || (false && true) || !(false && false)
	fmt.Println(expr)
}
