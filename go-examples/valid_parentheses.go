package main

import "fmt"

func isValid(s string) bool {
	var res bool
	if 1 > len(s) || len(s) >104 {
		return res
	}
	for i := range s {
		if i > 0 {
			previous := s[i-1]
			switch s[i] {
			case ')':
				if previous != '(' {
					return res
				}
			case ']':
				if previous != '[' {
					return res
				}
			case '}':
				if previous != '{' && previous != ']' {
					return res
				}
			}
		}
	}

	return true
}

func main() {
	//s := "()[]{}"
	s := "{[]}"
	valid := isValid(s)
	fmt.Println(valid)
}

