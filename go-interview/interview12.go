package main

import "fmt"

type People struct{}

func (p *People) ShowA() {
	fmt.Println("People ShowA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("People ShowB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("Teacher ShowB")
}

func main() {
	// keywords in golang
	// func
	// struct
	// defer

	a := -5
	b := +5
	fmt.Printf("a: %+d, b: %+d\n", a, b) // a: -5, b: +5

	t := Teacher{}
	t.ShowB() // Teacher ShowB
}
