package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		//value := val
		//m[key] = &value
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

//3 -> 3
//0 -> 3
//1 -> 3
//2 -> 3
// need to copy other memory
