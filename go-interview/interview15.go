package main

import "fmt"

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

func main() {
	var s1 []int //nil
	//var s2 = []int{}	//empty slice
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}

	i := 65
	fmt.Println(string(i)) // UTF-8 code, 65 decimal is A
	fmt.Printf("%d\n", i)  // 65

	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA()) // 13
	fmt.Println(b.ShowB()) // 23
}
