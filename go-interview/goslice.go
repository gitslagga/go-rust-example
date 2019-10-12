package main

import "fmt"

func main() {
	s := []int{5}
	fmt.Println(cap(s)) //1

	s = append(s, 7)
	fmt.Println(cap(s)) //2

	s = append(s, 9)
	fmt.Println(cap(s)) //4

	x := append(s, 11)
	y := append(s, 12)

	fmt.Println(s, x, y) //[5 7 9] [5 7 9 12] [5 7 9 12]
}
