package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		fmt.Println("Try to recover the panic.")
		if p := recover(); p != nil {
			fmt.Printf("Recoverd the painc(%#v).\n", p)
		}
	}()
	var mutex sync.Mutex
	fmt.Println("Lock the lock.")
	mutex.Lock()
	fmt.Println("The lock is locked.")
	fmt.Println("Unlock the lock.")
	mutex.Unlock()
	fmt.Println("The lock is unlocked.")
	fmt.Println("Unlock the lock again.")
	mutex.Unlock() //can't recoverd still get unlock error
}
