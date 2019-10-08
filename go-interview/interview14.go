package main

import "fmt"

func incr(p *int) int {
	*p++
	return *p
}

func add(args ...int) int {
	var sum int
	for _, arg := range args {
		sum += arg
	}
	fmt.Println(sum)
	return sum
}

func main() {
	str := "hello"
	str[0] = 'x'
	fmt.Println(str) // cannot assign to str[0]

	p := 1
	incr(&p)
	fmt.Println(p) // 2

	add(1, 2)              // 3
	add(1, 2, 3)           // 6
	add([]int{1, 3, 7}...) // 11
}
