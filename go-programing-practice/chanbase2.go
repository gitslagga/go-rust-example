package main

import (
	"fmt"
	"time"
)

var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2)
	go send(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a aync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "receiver")
		} else {
			break
		}
	}

	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Send:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Send a async signal. [sender]")
		}
	}

	fmt.Println("Wait 2 second... [sender]")
	time.Sleep(2 * time.Second)
	close(strChan)
	syncChan2 <- struct{}{}
}

//Send: a [sender]
//Send: b [sender]
//Send: c [sender]
//Send a async signal. [sender]
//Received a aync signal and wait a second... [receiver]
//Send: d [sender]
//Received: a receiver
//Received: b receiver
//Received: c receiver
//Received: d receiver
//Wait 2 second... [sender]
//Stopped. [receiver]
