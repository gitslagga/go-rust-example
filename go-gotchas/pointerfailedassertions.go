package main

import "fmt"

func main() {
	// Failed Type Assertions
	var data interface{} = "great"

	// Incorrect
	//if data, ok := data.(int); ok {
	//	fmt.Println("[is an int] value =>",data)
	//} else {
	//	fmt.Println("[not an int] value =>",data)
	//	//prints: [not an int] value => 0 (not "great")
	//}

	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data)
		//prints: [not an int] value => great (as expected)
	}
}
