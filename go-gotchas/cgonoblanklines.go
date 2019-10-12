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

	// One of the first gotchas with Cgo is the location of the cgo comments above the import "C" statement.
}

// cgonoblanklines.go:16:2: could not determine kind of name for C.free
