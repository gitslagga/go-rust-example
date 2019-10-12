package main

/*
#include <stdio.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

func main() {
	// Can't Call C Functions with Variable Arguments
	cstr := C.CString("go")
	C.printf("%s\n", cstr) //not ok
	C.free(unsafe.Pointer(cstr))

	// You can't call C functions with variable arguments directly.
	// You have to wrap your variadic C functions in functions with a known number of parameters.
}

// cgocallcfunctions.go:16:2: unexpected type: ...
