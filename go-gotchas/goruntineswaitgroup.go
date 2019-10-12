package main

import (
	"fmt"
	"sync"
)

func main() {
	// App Exits With Active Goroutines
	var wg sync.WaitGroup
	done := make(chan struct{})
	wq := make(chan interface{})
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go doitW(i, wq, done, &wg)
	}

	for i := 0; i < workerCount; i++ {
		wq <- i
	}

	// done<- struct{}{}
	// done<- struct{}{}
	close(done)
	wg.Wait()
	fmt.Println("All done!")
}

func doitW(workerId int, wq <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Printf("[%v] is running\n", workerId)
	defer wg.Done()
	for {
		select {
		case m := <-wq:
			fmt.Printf("[%v] m => %v\n", workerId, m)
		case <-done:
			fmt.Printf("[%v] is done\n", workerId)
			return
		}
	}
}

// [0] is running
// [0] is done
// [1] is running
// [1] is done
// All done!
