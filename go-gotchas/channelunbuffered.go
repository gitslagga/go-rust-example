package main

import "fmt"

func main() {
	// Sending to an Unbuffered Channel Returns As Soon As the Target Receiver Is Ready
	ch := make(chan string)

	go func() {
		for m := range ch {
			fmt.Println("processed:", m)
		}
	}()

	ch <- "cmd.1"
	ch <- "cmd.2" //won't be processed
}

// processed: cmd.1
// processed: cmd.2
