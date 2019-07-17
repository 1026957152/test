package com

import (
	"fmt"
	"syscall"
	"test/src/hikvision"

	"unsafe"
)

// LedDriver represents a digital Led
type RfidDriver struct {
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
func NewRfidDriver(a DigitalWriter, pin string) *RfidDriver {
	l := &RfidDriver{
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

	return l
}

// Start implements the Driver interface
func (l *RfidDriver) Start() (err error) { return }

// Halt implements the Driver interface
func (l *RfidDriver) Halt() (err error) { return }

// Name returns the LedDrivers name
func (l *RfidDriver) Name() string { return l.name }

// SetName sets the LedDrivers name
func (l *RfidDriver) SetName(n string) { l.name = n }

// Pin returns the LedDrivers name
func (l *RfidDriver) Pin() string { return l.pin }

// Connection returns the LedDrivers Connection
func (l *RfidDriver) Connection() hikvision.Connection {
	return l.connection.(hikvision.Connection)
}

// State return true if the led is On and false if the led is Off
func (l *RfidDriver) State() bool {
	return l.high
}

// On sets the led to a high state.
func (l *RfidDriver) On() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 1); err != nil {
		return
	}
	l.high = true
	return
}

// Off sets the led to a low state.
func (l *RfidDriver) Off() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 0); err != nil {
		return
	}
	l.high = false
	return
}

// Toggle sets the led to the opposite of it's current state
func (l *RfidDriver) Toggle() (err error) {
	if l.State() {
		err = l.Off()
	} else {
		err = l.On()
	}
	return
}

// Brightness sets the led to the specified level of brightness
func (l *RfidDriver) Brightness(level byte) (err error) {
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
