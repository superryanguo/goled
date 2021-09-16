package main

// #include "hello.h"
// #include <stdlib.h>
import "C"
import "unsafe"

func main() {
	cs := C.CString("Hello, World with cpp+go\n")
	defer C.free(unsafe.Pointer(cs))
	C.SayHello(cs)
}
