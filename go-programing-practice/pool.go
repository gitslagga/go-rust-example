package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func main() {
	//forbidden garbage collection
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}
	pool := sync.Pool{New: newFunc}

	//new field action
	v1 := pool.Get()
	fmt.Printf("Value 1: %v\n", v1)

	pool.Put(10)
	pool.Put(11)
	pool.Put(12)
	v2 := pool.Get()
	fmt.Printf("Value 2: %v\n", v2)

	//allow garbage collection
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("Value 3: %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("Value 4: %v\n", v4)
}

//Value 1: 1
//Value 2: 10
//Value 3: 2
//Value 4: <nil>
