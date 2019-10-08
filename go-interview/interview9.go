package main

import "fmt"

func world(num ...int) {
	num[0] = 18
}

func main() {
	// channel usage
	// write to channel need variable, when `ch <-` is wrong

	type person struct {
		name string
	}
	var m map[person]int
	p := person{name: "mike"}
	fmt.Println(m[p]) // 0

	i := []int{7, 8, 9}
	world(i...)
	fmt.Println(i) // [18 8 9]
}
