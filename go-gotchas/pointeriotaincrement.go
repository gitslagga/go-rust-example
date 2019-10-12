package main

import (
	"fmt"
)

const (
	azero = iota
	aone  = iota
)

const (
	info  = "processing"
	bzero = iota
	bone  = iota
)

func main() {
	// The First Use of iota Doesn't Always Start with Zero
	fmt.Println(azero, aone) //prints: 0 1
	fmt.Println(bzero, bone) //prints: 1 2
}
