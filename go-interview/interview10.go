package main

import (
	"fmt"
)

func main() {
	a := 5
	b := 8.1
	fmt.Println(a + b) // invalid operation: a + b (mismatched types int and float64)

	x := [5]int{1, 2, 3, 4, 5}
	t := x[3:4:4]
	fmt.Println(t) // [4]

	y := [2]int{5, 6}
	z := [3]int{5, 6}
	if y == z { // invalid operation: y == z (mismatched types [2]int and [3]int)
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
