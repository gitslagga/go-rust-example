package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)	//	[0 0 0 0 0 1 2 3]


	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)	//	[1 2 3 4]
}


func funcMui(x, y int) (sum int, error) {	// func has both named or unnamed return parameters
	return x + y, nil
}