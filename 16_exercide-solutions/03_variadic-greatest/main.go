package main

import "fmt"

func main() {
	greatest := findGreatest(23, 45, 67, 34, 12)

	fmt.Println("The greatest number is ", greatest)
}

func findGreatest(n ...int) int {
	var great int

	for _, v := range n {
		if v > great {
			great = v
		}
	}

	return great
}

/*	from solution
FYI
For your code to also work with only negative numbers such as
greatest := max(-200, -700)
include this as your range statement
for i, v := range numbers {
	if v > largest || i == 0 {
		largest = v
	}
}
What does that code do?
The first time through the range loop
the index, i, will be zero
so largest will be set to the first number
Originally largest is set to the zero value for an int, which is zero
Zero would be greater than any negative number
if you only have negative numbers
you need largest to be something less than zero
Thanks to Ricardo G for this code improvement!
*/
