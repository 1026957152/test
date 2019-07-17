package main

/*
#cgo CFLAGS: -I.


#cgo windows LDFLAGS: -lvideo
#cgo windows CFLAGS: -DWINDOWS
#cgo linux LDFLAGS: -lexample
#cgo LDFLAGS: -L"${SRCDIR}"
#cgo CFLAGS: -I"${SRCDIR}"

#include "video.h"
*/
import "C"
import "fmt"
import (
	"test/src/drivers"
	"test/src/drivers/com"
	"test/src/hikvision"
	"test/src/platforms/hik2"
	"test/src/platforms/rfid"
	"test/src/platforms/serial"
	"time"
)

func foo() {
	handle := C.dlopen(C.CString("libfoo.so"), C.RTLD_LAZY)
	bar := C.dlsym(handle, C.CString("bar"))
	fmt.Printf("bar is at %p\n", bar)
}

func main() {

	// 调用动态库函数fun1
	cmd := C.CString("ffmpeg -i ./xxx/*.png ./xxx/yyy.mp4")
	C.exeFFmpegCmd(&cmd)
	/*	// 调用动态库函数fun2
		C.fun2(C.int(4))
		// 调用动态库函数fun3
		var pointer unsafe.Pointer
		ret := C.fun3(&pointer)
		fmt.Println(int(ret))*/

	//C.puts(C.CString("你好，Cgo\n"))
	//C.SayHello(C.CString("你好，Cgo\n"))
	if true {
		return
	}

	firmataAdaptor := hik2.NewAdaptor("/dev/ttyACM0", "admin", "admin", 8000)
	led := hik.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		hikvision.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := hikvision.NewRobot("bot",
		[]hikvision.Connection{firmataAdaptor},
		[]hikvision.Device{led},
		work,
	)

	robot.Start()
}

func main() {

	firmataAdaptor := rfid.NewAdaptor("/dev/ttyACM0", "admin", "admin", 8000)
	led := com.NewRfidDriver(firmataAdaptor, "13")

	work := func() {
		hikvision.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := hikvision.NewRobot("bot",
		[]hikvision.Connection{firmataAdaptor},
		[]hikvision.Device{led},
		work,
	)

	robot.Start()
}

func mainZte() {

	firmataAdaptor := serial.NewAdaptor("/dev/ttyACM0", "admin", "admin", 8000)
	led := com.NewZteDriver(firmataAdaptor, "13")

	work := func() {
		hikvision.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := hikvision.NewRobot("bot",
		[]hikvision.Connection{firmataAdaptor},
		[]hikvision.Device{led},
		work,
	)

	robot.Start()
}
