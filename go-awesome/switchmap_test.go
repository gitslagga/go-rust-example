package go_awesome

import "testing"

func BenchmarkSwitch(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OrderStateSwitch(0)
		OrderStateSwitch(1)
		OrderStateSwitch(2)
		OrderStateSwitch(3)
	}
}

func BenchmarkMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		OrderStateMap(0)
		OrderStateMap(1)
		OrderStateMap(2)
		OrderStateMap(3)
	}
}
