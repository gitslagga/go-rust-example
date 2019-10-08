package main

import "fmt"

func main() {
	// how string, int variables connection
	// str := "abc" + "123"
	// fmt.Sprintf("abc%d", 123)

	const (
		x = iota
		_
		y
		z = "zz"
		k
		p = iota
	)
	fmt.Println(x, y, z, k, p)
	//0 2 zz zz 5

	// type chan, func, interface, map, slice, error have the default value is nil
	// var x interface{} = nil
	// var x error = nil
}
