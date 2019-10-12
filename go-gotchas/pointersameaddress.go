package main

import "fmt"

type data struct {
}

func main() {
	// Same Address for Different Zero-sized Variables
	a := &data{}
	b := &data{}

	if a == b {
		fmt.Printf("same address - a=%p b=%p\n", a, b)
		//prints: same address - a=0x1953e4 b=0x1953e4
	}
}
