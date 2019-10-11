package main

import "fmt"

func main() {
	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(x)

	fmt.Println(x) //prints [1 2 3] (not ok if you need [7 2 3])

	// use array
	y := [3]int{1, 2, 3}

	func(arr *[3]int) {
		(*arr)[0] = 7
		fmt.Println(arr) //prints &[7 2 3]
	}(&y)

	fmt.Println(y) //prints [7 2 3]

	// use slice
	z := []int{1, 2, 3}

	func(arr []int) {
		arr[0] = 7
		fmt.Println(arr) //prints [7 2 3]
	}(z)

	fmt.Println(z) //prints [7 2 3]
}
