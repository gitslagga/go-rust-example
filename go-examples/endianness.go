package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

func IsLittleEndian() bool {
	var value uint32 = 1
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
		panic(err)
	}

	fmt.Println(buf.Bytes())
	if (*pb) != buf.Bytes()[0] {
		return false
	}
	return true
}

func IsBigEndian() bool {
	var value uint32 = 1
	pointer := unsafe.Pointer(&value)
	pb := (*byte)(pointer)

	buf := new(bytes.Buffer)
	if err := binary.Write(buf, binary.BigEndian, value); err != nil {
		panic(err)
	}

	fmt.Println(buf.Bytes())
	if (*pb) != buf.Bytes()[0] {
		return false
	}
	return true
}

func BigEndianAndLittleEndian() {
	var value uint32 = 10
	by := make([]byte, 4)
	binary.BigEndian.PutUint32(by, value)
	big := binary.BigEndian.Uint32(by)
	fmt.Println("convert to big endian:", by)
	fmt.Printf("binary is: %b\n", big)
	fmt.Println("big endian result:", big)

	little := binary.LittleEndian.Uint32(by)
	fmt.Println()
	fmt.Printf("binary is: %b\n", little)
	fmt.Println("little endian result：", little)
}

func main() {
	fmt.Println(IsLittleEndian())
	fmt.Println()

	fmt.Println(IsBigEndian())
	fmt.Println()

	BigEndianAndLittleEndian()
}
