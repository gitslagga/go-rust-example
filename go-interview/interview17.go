package main

import "fmt"

func increateA() int {
	var i int
	defer func() {
		i++
	}()
	return i
}

func increateB() (r int) {
	defer func() {
		r++
	}()
	return 0
}

type X interface {
	ShowA() int
}

type Y interface {
	ShowB() int
}

type Life struct {
	i int
}

func (w Life) ShowA() int {
	return w.i + 10
}

func (w Life) ShowB() int {
	return w.i + 20
}

func main() {
	// x has statement, y has not statement, judgement program
	//1. x, _ := f()
	//2. x, _ = f()		right
	//3. x, y := f()	right
	//4. x, y = f()

	fmt.Println(increateA()) // 0
	fmt.Println(increateB()) // 1

	var x X = Life{3}
	s := x.(Life)
	fmt.Println(s.ShowA()) // 13
	fmt.Println(s.ShowB()) // 23
}
