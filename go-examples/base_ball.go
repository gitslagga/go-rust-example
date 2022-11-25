package main

import (
	"fmt"
	"strconv"
)

func calPoints(ops []string) int {
	var res int
	if 1 > len(ops) || len(ops) > 1000 {
		return res
	}

	var temp []int
	for i := range ops {

		switch ops[i] {
		case "C":
			if len(temp) > 0 {
				temp = temp[:len(temp)-1]
			}
		case "D":
			if len(temp) > 0 {
				mul := temp[len(temp)-1]*2
				temp = append(temp, mul)
			}
		case "+":
			if len(temp) > 1 {
				sum := temp[len(temp)-2] + temp[len(temp)-1]
				temp = append(temp, sum)
			}
		default:
			num, _ := strconv.Atoi(ops[i])
			temp = append(temp, num)
		}
	}
	for i := range temp {
		res += temp[i]
	}
	return res
}

func main() {
	s := []string{"5", "2", "C", "D", "+"}
	total := calPoints(s)
	fmt.Println(total)
}
