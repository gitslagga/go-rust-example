package main

import "fmt"

type data struct {
	name string
}

func main() {
	// Updating Map Value Fields
	// Incorrect
	//m := map[string]data {"x":{"one"}}
	//m["x"].name = "two" //error

	m := map[string]data{"x": {"one"}}
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v\n", m) //prints: map[x:{two}]

	mp := map[string]*data{"x": {"one"}}
	mp["x"].name = "two" //ok
	fmt.Println(mp["x"]) //prints: &{two}
}
