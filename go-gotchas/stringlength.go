package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	data := "♥"
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 1

	data = "é"
	fmt.Println(len(data))                    //prints: 3
	fmt.Println(utf8.RuneCountInString(data)) //prints: 2

	data = "我a "
	fmt.Println(len(data))                    //prints: 5
	fmt.Println(utf8.RuneCountInString(data)) //prints: 3
}
