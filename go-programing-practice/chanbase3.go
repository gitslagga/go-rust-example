package main

import (
	"fmt"
	"time"
)

func main() {
	var strChan = make(chan string, 3)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive3(strChan, syncChan1, syncChan2)
	go send3(strChan, syncChan1, syncChan2)
	<-syncChan2
	<-syncChan2
}

func receive3(strChan <-chan string, syncChan1 <-chan struct{}, syncChan2 chan<- struct{}) {
	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for elem := range strChan {
		fmt.Println("Received:", elem, "[receiver]")
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send3(strChan chan<- string, syncChan1 chan<- struct{}, syncChan2 chan<- struct{}) {
	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Send:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Send a sync a signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}

//Send: a [sender]
//Send: b [sender]
//Send: c [sender]
//Send a sync a signal. [sender]
//Received a sync signal and wait a second... [receiver]
//Received: a [receiver]
//Received: b [receiver]
//Received: c [receiver]
//Received: d [receiver]
//Send: d [sender]
//Wait 2 seconds... [sender]
//Stopped. [receiver]
