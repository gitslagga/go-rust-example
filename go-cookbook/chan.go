package main

import (
	"fmt"
	"time"
)

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("Send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("Receive:", v)
	}
}

func main() {
	//1: first		result: 1 2
	//ch :=make(chan int,1)
	//ch <- 1
	//go func() {
	//	v := <-ch
	//	fmt.Println(v)
	//}()
	//time.Sleep(1 * time.Second)
	//fmt.Println("2")

	//2: second		result: 1 2
	//ch := make(chan int)
	//go func() {
	//	v := <-ch
	//	fmt.Println(v)
	//}()
	//ch <- 1
	//fmt.Println(2)

	//3, third		result: 0 0 1 1 2 2 ......
	//ch := make(chan int)
	//go produce(ch)
	//go consumer(ch)
	//time.Sleep(1 * time.Second)

	//4, four		result: 0 1 2 3 4 5 ......
	ch := make(chan int, 10)
	go produce(ch)
	go consumer(ch)
	time.Sleep(time.Second)
}

