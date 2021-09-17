package main

//#cgo CFLAGS: -I ./myclib/
//#cgo LDFLAGS: -L ./myclib -lmyclib
// #include <stdlib.h>
// #include "./myclib/mytest.h"
import "C"
import (
	"fmt"
	"unsafe"
)

func Printc(s string) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.Printt(cs)
	return 0
}

func main() {
	fmt.Println(int(C.random()))
	Printc("hello cgo")
}
