package main

import (
	"fmt"
	"runtime"
)

func main() {
	// Preemptive Scheduling
	done := false

	go func() {
		done = true
	}()

	for !done {
		fmt.Println("not done!") //not inlined
	}

	for !done {
		runtime.Gosched()
	}

	fmt.Println("done!")
}

// not done!
// done!
