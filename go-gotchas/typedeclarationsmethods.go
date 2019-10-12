package main

import "sync"

type myLocker1 struct {
	sync.Mutex
}

type myLocker2 sync.Locker

func main() {
	// Type Declarations and Methods
	var lock1 myLocker1
	lock1.Lock()   //ok
	lock1.Unlock() //ok

	var lock2 myLocker2 = new(sync.Mutex)
	lock2.Lock()   //ok
	lock2.Unlock() //ok
}
