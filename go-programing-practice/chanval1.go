package main

import (
	"fmt"
	"time"
)

func main() {
	var mapChan = make(chan map[string]int, 1)

	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if elem, ok := <-mapChan; ok {
				elem["count"]++
			} else {
				break
			}
		}
		fmt.Println("Stopped. [receiver]")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- countMap
			time.Sleep(time.Millisecond)
			fmt.Printf("The count map: %v. [sender]\n", countMap)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan
}

//The count map: map[count:1]. [sender]
//The count map: map[count:2]. [sender]
//The count map: map[count:3]. [sender]
//The count map: map[count:4]. [sender]
//The count map: map[count:5]. [sender]
//Stopped. [receiver]
