package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Printf("Present time: %v.\n", time.Now())
	expirationTime := <-timer.C
	fmt.Printf("Expiration time: %v.\n", expirationTime)
	fmt.Printf("Stop timer: %v.\n", timer.Stop())
}

//Present time: 2019-09-07 17:37:10.8010465 +0800 CST m=+0.002992701.
//Expiration time: 2019-09-07 17:37:12.8090607 +0800 CST m=+2.011006801.
//Stop timer: false.
