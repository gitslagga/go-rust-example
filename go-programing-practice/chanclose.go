package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		for {
			if elem, ok := <-dataChan; ok {
				fmt.Printf("Received: %d [receiver]\n", elem)
			} else {
				break
			}
		}
		fmt.Println("Done. [receiver]")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("Send: %d [Sender]\n", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		fmt.Println("Done. [sender]")
		syncChan2 <- struct{}{}
	}()
	<-syncChan2
	<-syncChan2
}

//Send: 0 [Sender]
//Send: 1 [Sender]
//Send: 2 [Sender]
//Send: 3 [Sender]
//Send: 4 [Sender]
//Done. [sender]
//Received: 0 [receiver]
//Received: 1 [receiver]
//Received: 2 [receiver]
//Received: 3 [receiver]
//Received: 4 [receiver]
//Done. [receiver]
