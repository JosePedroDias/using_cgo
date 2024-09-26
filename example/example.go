package example

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L./ -lexample
#include "example.h"
*/
import "C"

// Wrap the C function in Go
func Add(a, b int) int {
	return int(C.add(C.int(a), C.int(b)))
}
