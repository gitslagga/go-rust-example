package main

import "fmt"

func hello() []int {
	return nil
}

func GetValue() int {
	return 1
}

func main() {
	// init func description
	// 1, a package have more than one init func
	// 2, when program build, build dependency package init func first, then build main package init func

	h := hello // this is hello func not hello func return variable
	if h == nil {
		fmt.Println("h is nil")
	} else {
		fmt.Println("h is not nil")
	}
	// h is not nil

	i := GetValue()
	switch i.(type) { // cannot type switch on non-interface value i (type int)
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}
}
