package main

//#cgo CFLAGS: -I ../oledc/
//#cgo LDFLAGS: -L ../oledc -loledc -lwiringPi
// #include <stdlib.h>
// #include "../oledc/oled.h"
// #include "../oledc/oled_fonts.h"
import "C"
import (
	"fmt"
	"image"
	"image/color"
	"os"
	"time"
	"unsafe"

	"gocv.io/x/gocv"
)

func OledInit() {
	C.oled_init()
}

func OledShow(s string) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	C.ssd1306_clearDisplay()
	C.ssd1306_drawText(C.int(0), C.int(8), cs)
	C.ssd1306_display()
}

func main() {
	fmt.Println("Begin the show...", int(C.random()))
	OledInit()
	OledShow("Using the OLED show")
	time.Sleep(500 * time.Microsecond)

	if len(os.Args) < 3 {
		fmt.Println("How to run:\n\tfacedetect [camera ID] [classifier XML file]")
		return
	}

	// parse args
	deviceID := os.Args[1]
	xmlFile := os.Args[2]

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error opening video capture device: %v\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Face Detect")
	defer window.Close()

	// prepare image matrix
	img := gocv.NewMat()
	defer img.Close()

	// color for the rect when faces detected
	blue := color.RGBA{0, 0, 255, 0}

	// load classifier to recognize faces
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()

	if !classifier.Load(xmlFile) {
		fmt.Printf("Error reading cascade file: %v\n", xmlFile)
		return
	}

	fmt.Printf("Start reading device: %v\n", deviceID)
	c := make(chan string, 10)
	defer close(c)
	go func() {
		for {
			s := <-c
			fmt.Printf(s)
			OledShow(s)
			time.Sleep(500 * time.Microsecond)
		}

	}()
	for {
		if ok := webcam.Read(&img); !ok {
			fmt.Printf("Device closed: %v\n", deviceID)
			return
		}
		if img.Empty() {
			continue
		}

		// detect faces
		rects := classifier.DetectMultiScale(img)
		s := fmt.Sprintf("found %d faces\n", len(rects))
		c <- s

		// draw a rectangle around each face on the original image,
		// along with text identifing as "Human"
		for _, r := range rects {
			gocv.Rectangle(&img, r, blue, 3)

			size := gocv.GetTextSize("Human", gocv.FontHersheyPlain, 1.2, 2)
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, "Human", pt, gocv.FontHersheyPlain, 1.2, blue, 2)
		}

		// show the image in the window, and wait 1 millisecond
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}

}
