package main

import "fmt"

func main() {
	// Breaking Out of "for switch" and "for select" Code Blocks
loop:
	for {
		switch {
		case true:
			fmt.Println("breaking out...")
			break loop
		}
	}
	fmt.Println("out!")
}
