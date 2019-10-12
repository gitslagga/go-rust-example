package main

import (
	"fmt"
	"time"
)

func main() {
	// Sending to an Closed Channel Causes a Panic
	ch := make(chan int)
	for i := 0; i < 3; i++ {
		go func(idx int) {
			ch <- (idx + 1) * 2
		}(i)
	}

	//get the first result
	fmt.Println(<-ch)
	close(ch) //not ok (you still have other senders)
	//do other work
	time.Sleep(2 * time.Second)
}

// 2
// panic: send on closed channel
