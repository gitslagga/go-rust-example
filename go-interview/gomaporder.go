package main

import (
	"fmt"
	"sort"
)

func main() {
	m := make(map[string]string)
	m["hello"] = "echo hello"
	m["world"] = "echo world"
	m["go"] = "echo go"
	m["is"] = "echo is"
	m["cool"] = "echo cool"

	sortedKeys := make([]string, 0)
	for k, _ := range m {
		sortedKeys = append(sortedKeys, k)
	}

	// sort 'string' key in increasing order
	sort.Strings(sortedKeys)

	for _, k := range sortedKeys {
		fmt.Printf("k=%v, v=%v\n", k, m[k])
	}
}

// k=cool, v=echo cool
// k=go, v=echo go
// k=hello, v=echo hello
// k=is, v=echo is
// k=world, v=echo world
