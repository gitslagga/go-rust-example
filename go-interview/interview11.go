package main

import "fmt"

func main() {
	// About the type of use of the cap function
	// 1, array
	// 2, slice
	// 3, channel

	var i interface{}
	if i == nil {
		fmt.Println("nil") // nil
		return
	}
	fmt.Println("not nil")

	s := make(map[string]int)
	delete(s, "h")
	fmt.Println(s["h"]) // 0
}
