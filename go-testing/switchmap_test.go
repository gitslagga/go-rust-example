package go_testing

import "testing"

const (
	UnPay = iota
	HadPay
	Delivery
	Finish
)

var orderState = map[int]string{
	UnPay:    "未支付",
	HadPay:   "已支付",
	Delivery: "配送中",
	Finish:   "已完成",
}

// map 实现
func OrderStateMap(state int) string {
	return orderState[state]
}

// switch 实现
func OrderStateSwitch(state int) string {
	var stateDesc = ""

	switch state {
	case UnPay:
		stateDesc = "未支付"
	case HadPay:
		stateDesc = "已支付"
	case Delivery:
		stateDesc = "配送中"
	case Finish:
		stateDesc = "已完成"
	}

	return stateDesc
}

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

// go test -bench=.
