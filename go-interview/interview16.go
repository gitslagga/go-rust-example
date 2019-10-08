package main

import "fmt"

func main() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println(len(a), len(b), len(c)) // 0 2 1
	fmt.Println(cap(a), cap(b), cap(c)) // 3 3 2

	var m = make(map[string]int)
	m["a"] = 1
	if v, ok := m["b"]; ok {
		fmt.Println(v)
	}

	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB()) // a.ShowB undefined (type A has no field or method ShowB)
	fmt.Println(b.ShowA()) // b.ShowA undefined (type B has no field or method ShowA)
}
