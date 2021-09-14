package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func Printc(s string) int {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	//C.printf("c function use %s\n", cs)
	return 0
}

func main() {
	fmt.Println(int(C.random()))
	Printc("hello cgo")
}
