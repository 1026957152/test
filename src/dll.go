package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"
)

const (
	MB_OK                = 0x00000000
	MB_OKCANCEL          = 0x00000001
	MB_ABORTRETRYIGNORE  = 0x00000002
	MB_YESNOCANCEL       = 0x00000003
	MB_YESNO             = 0x00000004
	MB_RETRYCANCEL       = 0x00000005
	MB_CANCELTRYCONTINUE = 0x00000006
	MB_ICONHAND          = 0x00000010
	MB_ICONQUESTION      = 0x00000020
	MB_ICONEXCLAMATION   = 0x00000030
	MB_ICONASTERISK      = 0x00000040
	MB_USERICON          = 0x00000080
	MB_ICONWARNING       = MB_ICONEXCLAMATION
	MB_ICONERROR         = MB_ICONHAND
	MB_ICONINFORMATION   = MB_ICONASTERISK
	MB_ICONSTOP          = MB_ICONHAND

	MB_DEFBUTTON1 = 0x00000000
	MB_DEFBUTTON2 = 0x00000100
	MB_DEFBUTTON3 = 0x00000200
	MB_DEFBUTTON4 = 0x00000300
)

func abort(funcname string, err syscall.Errno) {
	panic(funcname + " failed: " + err.Error())
}

var (
	user32, _     = syscall.LoadLibrary("user32.dll")
	messageBox, _ = syscall.GetProcAddress(user32, "MessageBoxW")
)

func IntPtr(n int) uintptr {
	return uintptr(n)
}

func StrPtr(s string) uintptr {
	return uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(s)))
}

func MessageBox(caption, text string, style uintptr) (result int) {
	ret, _, callErr := syscall.Syscall9(messageBox,
		4,
		0,
		StrPtr(text),
		StrPtr(caption),
		style,
		0, 0, 0, 0, 0)
	if callErr != 0 {
		abort("Call MessageBox", callErr)
	}
	result = int(ret)
	return
}

//func GetModuleHandle() (handle uintptr) {
//    if ret, _, callErr := syscall.Syscall(getModuleHandle, 0, 0, 0, 0); callErr != 0 {
//        abort("Call GetModuleHandle", callErr)
//    } else {
//        handle = ret
//    }
//    return
//}

// windows下的第二种DLL方法调用
func ShowMessage2(title, text string) {
	user32 := syscall.NewLazyDLL("user32.dll")
	MessageBoxW := user32.NewProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(title), IntPtr(0))
}

// windows下的第三种DLL方法调用
func ShowMessage3(title, text string) {
	user32, _ := syscall.LoadDLL("user32.dll")
	MessageBoxW, _ := user32.FindProc("MessageBoxW")
	MessageBoxW.Call(IntPtr(0), StrPtr(text), StrPtr(title), IntPtr(0))
}

func main__() {
	defer syscall.FreeLibrary(user32)

	num := MessageBox("Done Title", "This test is Done.", MB_YESNOCANCEL)
	fmt.Printf("Get Retrun Value Before MessageBox Invoked: %d\n", num)
	ShowMessage2("windows下的另一种DLL方法调用", "HELLO !")

	ShowMessage3("windows下的第三种DLL方法调用", "lyslyslys !")

	time.Sleep(3 * time.Second)
}

func init() {
	fmt.Print("Starting Up\n")
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
func DllTestDef_add2() {
	DllTestDef := syscall.MustLoadDLL("ConsoleApplication2.dll")
	add := DllTestDef.MustFindProc("panny")

	fmt.Println("+++++++MustFindProc：", add, "+++++++")
	ret, _, err := add.Call(IntPtr(0))
	if err != nil {
		fmt.Println("DllTestDef的运算结果为:", ret)
	}
}

func main______() {
	h, err := syscall.LoadLibrary("ConsoleApplication2.dll")
	if err != nil {
		fmt.Println("没有找到 ConsoleApplication2")

		//	abort("LoadLibrary", err)
	}
	defer syscall.FreeLibrary(h)
	proc, err := syscall.GetProcAddress(h, "pannyp")
	if err != nil {
		fmt.Println("没有找到 pannyp ConsoleApplication2")
	}
	_, _, _ = syscall.Syscall(uintptr(proc), 0, 0, 0, 0)

	//fmt.Println("--------%s---------",r)

}
func main() {
	//defer syscall.FreeLibrary(user32)

	DllMessage3("windows下的另一种DLL方法调用", "HELLO !")
	main______()
	//DllTestDef_add2()
	time.Sleep(3 * time.Second)
}
