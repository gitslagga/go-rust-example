package main

/*
#include <stdlib.h>
*/

import "C"

import (
	"unsafe"
)

func main() {
	// No blank lines Between Import C and Cgo Comments
	cs := C.CString("my go string")
	C.free(unsafe.Pointer(cs))
}

//
