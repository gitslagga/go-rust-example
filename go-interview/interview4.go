package main

import "fmt"

func main() {
	list := new([]int)
	list = append(list, 1)	//cannot use 'list' (type *[]int) as type []Type
	fmt.Println(list)


	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s3 := append(s1, s2)	//cannot use 's2' (type []int) as type int
	fmt.Println(s3)


	var (
		size    := 1024		//excepted '=', got ':='
		maxSize = 2 * size
	)
	fmt.Println(size, maxSize)
}
