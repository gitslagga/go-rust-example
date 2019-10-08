package main

import "fmt"

func increateA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increateB() (r int) {
	defer func() {
		r++
	}()
	return 0
}

func main() {
	fmt.Println(increateA()) // 0
	fmt.Println(increateB()) // 1
}
