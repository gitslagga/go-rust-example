package main

import "fmt"

func main() {
	// "nil" Interfaces and "nil" Interfaces Values
	var data *byte
	var in interface{}

	fmt.Println(data, data == nil) //prints: <nil> true
	fmt.Println(in, in == nil)     //prints: <nil> true

	in = data
	fmt.Println(in, in == nil) //prints: <nil> false
	//'data' is 'nil', but 'in' is not 'nil'

	doit := func(arg int) interface{} {
		var result *struct{} = nil

		if arg > 0 {
			result = &struct{}{}
		} else {
			return nil //return an explicit 'nil'
		}

		return result // result is not 'nil', but its value is 'nil'
	}

	if res := doit(-1); res != nil {
		fmt.Println("good result:", res)
	} else {
		fmt.Println("bad result (res is nil)") //here as expected
	}
}
