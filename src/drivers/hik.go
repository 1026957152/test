package hik

import (
	"fmt"
	"syscall"
	"test/src/hikvision"
	"time"
	"unsafe"
)

// LedDriver represents a digital Led
type HikDriver struct {
	pin        string
	name       string
	connection DigitalWriter
	high       bool
	hikvision.Commander
}

// NewLedDriver return a new LedDriver given a DigitalWriter and pin.
//
// Adds the following API Commands:
//	"Brightness" - See LedDriver.Brightness
//	"Toggle" - See LedDriver.Toggle
//	"On" - See LedDriver.On
//	"Off" - See LedDriver.Off
func NewLedDriver(a DigitalWriter, pin string) *HikDriver {
	l := &HikDriver{
		name:       hikvision.DefaultName("LED"),
		pin:        pin,
		connection: a,
		high:       false,
		Commander:  hikvision.NewCommander(),
	}

	l.AddCommand("Brightness", func(params map[string]interface{}) interface{} {
		level := byte(params["level"].(float64))
		return l.Brightness(level)
	})

	l.AddCommand("Toggle", func(params map[string]interface{}) interface{} {
		return l.Toggle()
	})

	l.AddCommand("On", func(params map[string]interface{}) interface{} {
		return l.On()
	})

	l.AddCommand("Off", func(params map[string]interface{}) interface{} {
		return l.Off()
	})
	l.AddCommand("云台上", func(params map[string]interface{}) interface{} {
		level := byte(params["level"].(float64))
		return l.云台上(level)
	})

	return l
}

// Start implements the Driver interface
func (l *HikDriver) Start() (err error) { return }

// Halt implements the Driver interface
func (l *HikDriver) Halt() (err error) { return }

// Name returns the LedDrivers name
func (l *HikDriver) Name() string { return l.name }

// SetName sets the LedDrivers name
func (l *HikDriver) SetName(n string) { l.name = n }

// Pin returns the LedDrivers name
func (l *HikDriver) Pin() string { return l.pin }

// Connection returns the LedDrivers Connection
func (l *HikDriver) Connection() hikvision.Connection {
	return l.connection.(hikvision.Connection)
}

// State return true if the led is On and false if the led is Off
func (l *HikDriver) State() bool {
	return l.high
}

// On sets the led to a high state.
func (l *HikDriver) On() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 1); err != nil {
		return
	}
	l.high = true
	return
}

// Off sets the led to a low state.
func (l *HikDriver) Off() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 0); err != nil {
		return
	}
	l.high = false
	return
}

// Toggle sets the led to the opposite of it's current state
func (l *HikDriver) Toggle() (err error) {
	if l.State() {
		err = l.Off()
	} else {
		err = l.On()
	}
	return
}

// Brightness sets the led to the specified level of brightness
func (l *HikDriver) Brightness(level byte) (err error) {
	if writer, ok := l.connection.(PwmWriter); ok {
		return writer.PwmWrite(l.Pin(), level)
	}
	return ErrPwmWriteUnsupported
}

// Brightness sets the led to the specified level of brightness
func (l *HikDriver) 云台上(level byte) (err error) {

	PTZControlAll(lPlayHandle, TILT_UP, 1, iPTZSpeed)
	if writer, ok := l.connection.(PwmWriter); ok {
		return writer.PwmWrite(l.Pin(), level)
	}

	return ErrPwmWriteUnsupported
}

// windows下的第三种DLL方法调用
func DllMessage3(title, text string) {
	user32, err := syscall.LoadDLL("ConsoleApplication2.dll")
	if err != nil {
		fmt.Println("没有找到 ConsoleApplication2")

	}
	MessageBoxW, err := user32.FindProc("pannyp")
	if err != nil {
		fmt.Println("没有找到 pannyp ConsoleApplication2")

	}
	MessageBoxW.Call(IntPtr(0))
}

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

/*
case 0:        //云台上
PTZControlAll(lPlayHandle, TILT_UP, 1, iPTZSpeed);
break;
case 1:			//云台下
PTZControlAll(lPlayHandle, TILT_DOWN, 1, iPTZSpeed);
break;
case 2:			//云台左
PTZControlAll(lPlayHandle, PAN_LEFT, 1, iPTZSpeed);
break;
case 3:			//云台右
PTZControlAll(lPlayHandle, PAN_RIGHT, 1, iPTZSpeed);
break;
case 4:         //调焦左
PTZControlAll(lPlayHandle, ZOOM_IN, 1, iPTZSpeed);
break;
case 5:			//调焦右
PTZControlAll(lPlayHandle, ZOOM_OUT, 1, iPTZSpeed);
break;
case 6:			//聚焦左
PTZControlAll(lPlayHandle, FOCUS_NEAR, 1, iPTZSpeed);
break;
case 7:			//聚焦右
PTZControlAll(lPlayHandle, FOCUS_FAR, 1, iPTZSpeed);
break;
case 8:			//光圈左
PTZControlAll(lPlayHandle, IRIS_OPEN, 1, iPTZSpeed);
break;
case 9:			//光圈右
PTZControlAll(lPlayHandle, IRIS_CLOSE, 1, iPTZSpeed);
break;
case 10:        //左上
PTZControlAll(lPlayHandle, UP_LEFT, 1, iPTZSpeed);
break;
case 11:        //右上
PTZControlAll(lPlayHandle, UP_RIGHT, 1, iPTZSpeed);
break;
case 12:        //左下
PTZControlAll(lPlayHandle, DOWN_LEFT, 1, iPTZSpeed);
break;
case 13:        //右下
PTZControlAll(lPlayHandle, DOWN_RIGHT, 1, iPTZSpeed);





抓图：
*/
