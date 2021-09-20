package main

//#cgo CFLAGS: -I ./oledc/
//#cgo LDFLAGS: -L ./oledc -loledc -lwiringPi
// #include <stdlib.h>
// #include "./oledc/oled.h"
// #include "./oledc/oled_fonts.h"
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

func OledInit() {
	C.oled_init()
}

func OledShow(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.ssd1306_drawText(C.int(0), C.int(8), cs)
	C.ssd1306_display()
}

func main() {
	fmt.Println("Begin the show...", int(C.random()))
	OledInit()
	for {
		OledShow("Use the OLED SHOW")
		time.Sleep(500 * time.Microsecond)
	}

}
