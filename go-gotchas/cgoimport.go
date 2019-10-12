package main

/*
#include <stdlib.h>
*/
import (
	"C"
	"unsafe"
)

func main() {
	// Import C and Multiline Import Blocks
	cs := C.CString("my go string")
	C.free(unsafe.Pointer(cs))
}

//
