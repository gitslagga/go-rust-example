package main

import (
	"fmt"
	"strings"
)

func calculate(str, num string) int {
	if len(str) != 10 {
		return 0
	}
	if 1 > len(num) || len(num) > 104 {
		return 0
	}

	var pre int
	var next int
	var order = strings.Index(str, string(num[0]))
	for i := range num {
		if i > 0 {
			pre = strings.Index(str, string(num[i-1]))
			next = strings.Index(str, string(num[i]))

			if pre > next {
				order += pre - next
			} else {
				order += next - pre
			}
		}
	}
	return order
}

func main() {
	str := "0123456789"
	num := "210"

	fmt.Println(calculate(str, num))
}
