package main

import (
	"fmt"
	"time"
)

func main() {
	sendingInterval := time.Second
	receptionInterval := time.Second * 2
	intChan := make(chan int, 0)
	go func() {
		var ts0, ts1 int64
		for i := 1; i <= 5; i++ {
			intChan <- i
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Send:", i)
			} else {
				fmt.Printf("Send: %d [interval: %d s]\n", i, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(sendingInterval)
		}
		close(intChan)
	}()
	var ts0, ts1 int64

Loop:
	for {
		select {
		case v, ok := <-intChan:
			if !ok {
				break Loop
			}
			ts1 = time.Now().Unix()
			if ts0 == 0 {
				fmt.Println("Received:", v)
			} else {
				fmt.Printf("Received: %d [interval: %d s]\n", v, ts1-ts0)
			}
			ts0 = time.Now().Unix()
			time.Sleep(receptionInterval)
		}
	}
	fmt.Println("end.")
}

//Send: 1
//Received: 1
//Send: 2 [interval: 2 s]
//Received: 2 [interval: 2 s]
//Received: 3 [interval: 2 s]
//Send: 3 [interval: 2 s]
//Received: 4 [interval: 2 s]
//Send: 4 [interval: 2 s]
//Received: 5 [interval: 2 s]
//Send: 5 [interval: 2 s]
//end.
