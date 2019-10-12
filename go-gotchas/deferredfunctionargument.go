package main

import "fmt"

func main() {
	// Deferred Function Call Argument Evaluation
	var i int = 1

	defer fmt.Println("result =>", func(in *int) int { return *in * 2 }(&i)) // result => 2

	defer func(in *int) {
		fmt.Println("result =>", *in*2) // result => 4
	}(&i)
	i++
}
