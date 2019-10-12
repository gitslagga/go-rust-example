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

	// If you are using the import block format you can't import other packages in the same block.
}

// cgoimport.go:14:2: could not determine kind of name for C.free
