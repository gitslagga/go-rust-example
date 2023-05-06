package main

import (
	"fmt"
)

var s = []int{1, 2, 3, 4}

func main() {
	var count int

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(s); k++ {
				if s[i] != s[j] && s[j] != s[k] && s[i] != s[k] {
					count++
					fmt.Println(s[i], s[k], s[j])
				}
			}
		}
	}

	fmt.Println("arrangement", count)
	fmt.Println("combination", count/2)
}
