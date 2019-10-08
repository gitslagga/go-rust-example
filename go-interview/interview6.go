package main

import "fmt"

func main() {
	// get member variables by pointer variable
	// 1, p.name
	// 2, (*p).name

	type MyInt1 int
	type MyInt2 = int
	var i int = 0
	var i1 MyInt1 = i // cannot use i (type int) as type MyInt1 in assignment
	var i2 MyInt2 = i
	fmt.Println(i1, i2)

	s := []int{7, 8, 9}
	fmt.Println(s)
	ap(s)
	fmt.Println(s)
	app(s)
	fmt.Println(s)
	//[7 8 9]
	//[7 8 9]
	//[10 8 9]
}

func ap(s []int) {
	s = append(s, 10) // append causes the bottom of the array to reallocate memory
}

func app(s []int) {
	s[0] = 10
}
