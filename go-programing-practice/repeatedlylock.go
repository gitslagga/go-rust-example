package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	fmt.Println("Lock the lock. (main)")
	mutex.Lock()
	fmt.Println("The lock is locked. (main)")
	for i := 1; i <= 3; i++ {
		go func(i int) {
			fmt.Printf("Lock the lock. (g%d)\n", i)
			mutex.Lock()
			fmt.Printf("The lock is locked. (g%d)\n", i)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock. (main)")
	mutex.Unlock()
	fmt.Println("The lock is unlocked. (main)")
	time.Sleep(time.Second)
}

//Mutex lock to unlock, more to one
//Lock the lock. (main)
//The lock is locked. (main)
//Lock the lock. (g1)
//Lock the lock. (g2)
//Lock the lock. (g3)
//Unlock the lock. (main)
//The lock is unlocked. (main)
//The lock is locked. (g1)
