package go_testing

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"
)

var count uint64
var countAtomic uint64

func TestAtomic(t *testing.T) {
	for i := 0; i < 10000; i++ {
		go func() {
			count++
			atomic.AddUint64(&countAtomic, 1)
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println(count)
	fmt.Println(atomic.LoadUint64(&countAtomic))
}

func TestAtomicSwap(t *testing.T) {
	oldValue := atomic.LoadUint64(&countAtomic)
	newValue := oldValue + 1
	for i := 0; i < 10000; i++ {
		go func() {
			if atomic.CompareAndSwapUint64(&countAtomic, oldValue, newValue) {
				count++
			}
		}()
	}

	time.Sleep(1 * time.Second)

	fmt.Println(count)
	fmt.Println(atomic.LoadUint64(&countAtomic))
}
