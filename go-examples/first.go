package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Convert a non-negative integer num to its English words representation.

	// Example 1:

	// Input: num = 123
	// Output: "One Hundred Twenty Three"
	// Example 2:

	// Input: num = 12345
	// Output: "Twelve Thousand Three Hundred Forty Five"
	// Example 3:

	// Input: num = 1234567
	// Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
	// Example 4:

	// Input: num = 1234567891
	// Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

	// Constraints:

	// 0 <= num <= 231 - 1

	NumberConvert(123)
	NumberConvert(12345)
	NumberConvert(1234567891)
}

var oneBit = map[string]string{
	"1": "One",
	"2": "Two",
	"3": "Three",
	"4": "Four",
	"5": "Five",
	"6": "Six",
	"7": "Seven",
	"8": "Eight",
	"9": "Nine",
}

var tenSimBit = map[string]string{
	"1": "Eleven",
	"2": "Twelve",
	"3": "Thirteen",
	"4": "Fourteen",
	"5": "Fifteen",
	"6": "Sixteen",
	"7": "Seventeen",
	"8": "Eighteen",
	"9": "Nineteen",
}

var tenBit = map[string]string{
	"2": "Twenty",
	"3": "Thirty",
	"4": "Fourty",
	"5": "Fifty",
	"6": "Sixty",
	"7": "Seventy",
	"8": "Eighty",
	"9": "Ninety",
}

func NumberConvert(num int) (out string) {
	if 0 <= num && num <= int(math.Pow(2, 31))-1 {
		strNum := strconv.Itoa(num)
		strNum = NumberReverse(strNum)

		for key, val := range strNum {
			switch key {
			case 0:
				out = oneBit[string(val)]
			case 1:
				if string(val) == "1" {
					out = tenSimBit[strNum[0:1]]
				} else {
					out = tenBit[string(val)] + " " + out
				}
			case 2:
				out = oneBit[string(val)] + " Hundred " + out
			case 3:
				if strNum[4:5] != "1" {
					out = oneBit[string(val)] + " Thousand " + out
				}
			case 4:
				if string(val) == "1" {
					out = tenSimBit[strNum[3:4]] + " Thousand " + out
				} else {
					out = tenBit[string(val)] + " " + out
				}
			case 5:
				out = oneBit[string(val)] + " Hundred " + out
			case 6:
				if strNum[7:8] != "1" {
					out = oneBit[string(val)] + " Million " + out
				}
			case 7:
				if string(val) == "1" {
					out = tenSimBit[strNum[6:7]] + " Million " + out
				} else {
					out = tenBit[string(val)] + " " + out
				}
			case 8:
				out = oneBit[string(val)] + " Hundred " + out
			case 9:
				out = oneBit[string(val)] + " Billion " + out
			}
		}
	}

	fmt.Println(out)
	return
}

func NumberReverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
