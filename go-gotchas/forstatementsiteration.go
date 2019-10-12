package main

import (
	"fmt"
	"time"
)

type field struct {
	name string
}

func (p *field) print() {
	fmt.Println(p.name)
}

func main() {
	// Iteration Variables and Closures in "for" Statements
	//data := []string{"one","two","three"}

	//for _,v := range data {
	//	vcopy := v //
	//	go func() {
	//		fmt.Println(vcopy)
	//	}()
	//}

	//for _,v := range data {
	//	go func(in string) {
	//		fmt.Println(in)
	//	}(v)
	//}

	data := []*field{{"one"}, {"two"}, {"three"}}

	for _, v := range data {
		go v.print()
	}

	time.Sleep(3 * time.Second)
}

// three
// two
// one
