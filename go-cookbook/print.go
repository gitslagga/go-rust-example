package main

import "fmt"

func main() {
	x := 1
	p := &x         // p, of type *int, points to x
	fmt.Println(*p) // "1"

	// Coroutine example
	*p = 2          // equivalent to x = 2
	println(x)  	// "2"

}
