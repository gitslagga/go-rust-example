package main

import "fmt"

func main() {
	var d uint8 = 2
	fmt.Printf("%08b\n", ^d)

	var a uint8 = 0x82
	var b uint8 = 0x02
	fmt.Printf("%08b [A]\n", a)
	fmt.Printf("%08b [B]\n", b)

	fmt.Printf("%08b (NOT B)\n", ^b)
	fmt.Printf("%08b ^ %08b = %08b [B XOR 0xff]\n", b, 0xff, b^0xff)

	fmt.Printf("%08b ^ %08b = %08b [A XOR B]\n", a, b, a^b)
	fmt.Printf("%08b & %08b = %08b [A AND B]\n", a, b, a&b)
	fmt.Printf("%08b &^%08b = %08b [A 'AND NOT' B]\n", a, b, a&^b)
	fmt.Printf("%08b |%08b = %08b [A 'OR' B]\n", a, b, a|b)
	fmt.Printf("%08b&(^%08b)= %08b [A AND (NOT B)]\n", a, b, a&(^b))

	fmt.Println(0xff ^ b)
	fmt.Println(0 ^ b)
}

// 11111101
// 10000010 [A]
// 00000010 [B]
// 11111101 (NOT B)
// 00000010 ^ 11111111 = 11111101 [B XOR 0xff]
// 10000010 ^ 00000010 = 10000000 [A XOR B]
// 10000010 & 00000010 = 00000010 [A AND B]
// 10000010 &^00000010 = 10000000 [A 'AND NOT' B]
// 10000010 |00000010 = 10000010 [A 'OR' B]
// 10000010&(^00000010)= 10000000 [A AND (NOT B)]
// 253
// 2
