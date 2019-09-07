package main

import (
	"fmt"
	"time"
)

func main() {
	var strChan = make(chan string, 3)

	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go func() {
		<-syncChan1
		fmt.Println("Received a sync signal and wait a second... [receiver]")
		time.Sleep(time.Second)
		for {
			if elem, ok := <-strChan; ok {
				fmt.Println("Received:", elem, "[receiver]")
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for _, elem := range []string{"a", "b", "c", "d"} {
			strChan <- elem
			fmt.Println("Send:", elem, "[sender]")
			if elem == "c" {
				syncChan1 <- struct{}{}
				fmt.Println("Send a async signal. [sender]")
			}
		}

		fmt.Println("Wait 2 seconds... [sender]")
		time.Sleep(2 * time.Second)
		close(strChan)
		syncChan2 <- struct{}{}
	}()

	<-syncChan2
	<-syncChan2
}

//Send: a [sender]
//Send: b [sender]
//Send: c [sender]
//Send a async signal. [sender]
//Received a sync signal and wait a second... [receiver]
//Received: a [receiver]
//Received: b [receiver]
//Received: c [receiver]
//Received: d [receiver]
//Send: d [sender]
//Wait 2 seconds... [sender]
//Stopped. [receiver]
