package main

import "fmt"

func main() {
	// Recovering From a Panic
	defer func() {
		fmt.Println("recovered:", recover())
	}()

	panic("not good")
}

// recovered: not good
