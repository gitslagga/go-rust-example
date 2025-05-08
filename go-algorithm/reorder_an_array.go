package main

import "fmt"

func solution(x []int, y []int) []int {
	res := []int{}

	for k, v := range x {
		if y[k] >= k {
			res = append(res, v)
		} else {
			res = append(res[:y[k]], append([]int{v}, res[y[k]:]...)...)
		}
	}
	return res
}

func main() {
	// 1, 5, 2, 4, 3
	recorder := solution([]int{1, 2, 3, 4, 5}, []int{0, 1, 2, 2, 1})
	fmt.Println(recorder)
}
