package main

import (
	"bytes"
	"fmt"
)

func main() {
	var b1 []byte = nil
	b2 := []byte{}
	fmt.Println("b1 == b2:", bytes.Equal(b1, b2)) //prints: b1 == b2: true
}
