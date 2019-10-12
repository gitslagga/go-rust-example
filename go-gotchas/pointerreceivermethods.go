package main

import "fmt"

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}

type printer interface {
	print()
}

func main() {
	// Using Pointer Receiver Methods On Value Instances
	d1 := data{"one"}
	d1.print() //ok

	var in printer = data{"two"} //error
	in.print()

	m := map[string]data{"x": data{"three"}}
	m["x"].print() //error
}
