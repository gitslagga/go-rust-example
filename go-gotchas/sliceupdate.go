package main

import "fmt"

func sliceUpdate() {
	data := []int{1, 2, 3}
	for i, _ := range data {
		data[i] *= 10
	}

	fmt.Println("data:", data) //prints data: [10 20 30]
}

func structUpdate() {
	data := []*struct{ num int }{{1}, {2}, {3}}

	for _, v := range data {
		v.num *= 10
	}

	fmt.Println(data[0], data[1], data[2]) //prints &{10} &{20} &{30}
}

func main() {
	// Updating and Referencing Item Values in Slice, Array, and Map "range" Clauses
	sliceUpdate()
	structUpdate()
}
