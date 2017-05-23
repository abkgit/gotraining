package main

import "fmt"

func max(x int) int {  //lowercase func name, so not exported(visible outside package)
	return 42 + x
}

func main() {
	max := max(7)
	fmt.Println(max) // max is now the result, not the function
}

// don't do this; bad coding practice to shadow variables