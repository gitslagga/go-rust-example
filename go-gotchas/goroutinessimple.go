package main

import (
	"fmt"
	"time"
)

func main() {
	// App Exits With Active Goroutines
	workerCount := 2

	for i := 0; i < workerCount; i++ {
		go doitS(i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("all done!")
}

func doitS(workerId int) {
	fmt.Printf("[%v] is running\n", workerId)
	time.Sleep(3 * time.Second)
	fmt.Printf("[%v] is done\n", workerId)
}

// [1] is running
// [0] is running
// all done!
