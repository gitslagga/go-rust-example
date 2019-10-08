package main

import "fmt"

func print(i int) {
	fmt.Println(i)
}

func main() {
	// global string define
	// 1, var str string
	// 2, var str = ""

	i := 5
	defer print(i) // 5
	i = i + 10

	t := Teacher{}
	t.ShowA()
	// People ShowA
	// People ShowB
}
