package main

import "sync"

type myMutex sync.Mutex

type myLocker1 struct {
	sync.Mutex
}

type myLocker2 sync.Locker

func main() {
	// Type Declarations and Methods
	// Incorrect
	//var mtx myMutex
	//mtx.Lock() //error
	//mtx.Unlock() //error

	var lock1 myLocker1
	lock1.Lock()   //ok
	lock1.Unlock() //ok

	var lock2 myLocker2 = new(sync.Mutex)
	lock2.Lock()   //ok
	lock2.Unlock() //ok
}
