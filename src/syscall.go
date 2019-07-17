package main

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"
)

var kernel32 syscall.Handle

//初始化获取方法的引用
func init() {
	var err error
	kernel32, err = syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("获取方法应用错误")
	}

}

func getDriveNames() ([]string, error) {

	drives := []string{}

	LongPtr_DriveBuf := make([]byte, 256)

	getDrivesStringsEx, err := syscall.GetProcAddress(kernel32, "GetLogicalDriveStringsW")
	if err != nil {
		return nil, errors.New("call GetLogicalDriveStringsW fail")
	}

	//执行调用
	// 因为有2个参数，所以使用syscall就能放得下，最后的参数补0
	r, _, errno := syscall.Syscall(uintptr(getDrivesStringsEx), 2,
		uintptr(len(LongPtr_DriveBuf)),
		uintptr(unsafe.Pointer(&LongPtr_DriveBuf[0])), 0)

	if r != 0 {

		for _, v := range LongPtr_DriveBuf {
			if v < 65 || v > 90 {
				continue
			}
			//println(string(v))
			drives = append(drives, string(v)+":")
		}

	} else {
		return nil, errors.New(errno.Error())
	}

	return drives, nil
}

func getDiskGreeSpace(diskName string) {

	//将磁盘的名称转化为*UTF16
	diskNameUTF16Ptr, _ := syscall.UTF16PtrFromString(diskName)

	//使用长指针
	LongPtr_FreeBytesAvailable := int64(0)     //剩余空间
	LongPtr_TotalNumberOfBytes := int64(0)     //总空间
	LongPtr_TotalNumberOfFreeBytes := int64(0) //可用空间

	//获取方法的引用
	kernel32, err := syscall.LoadLibrary("kernel32.dll")
	if err != nil {
		panic("获取方法应用错误")
	}

	//释放方法引用
	defer syscall.FreeLibrary(kernel32)

	getDiskFreeSpaceEx, err := syscall.GetProcAddress(kernel32, "GetDiskFreeSpaceExW")
	if err != nil {
		panic("call GetZDiskFreeSpaceExW fail")
	}

	// uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("C:"))),
	//执行调用
	// 因为有四个参数，所以使用syscall6才能放得下，最后两个参数补0
	r, _, errno := syscall.Syscall6(uintptr(getDiskFreeSpaceEx), 4,
		uintptr(unsafe.Pointer(diskNameUTF16Ptr)),
		uintptr(unsafe.Pointer(&LongPtr_FreeBytesAvailable)),
		uintptr(unsafe.Pointer(&LongPtr_TotalNumberOfBytes)),
		uintptr(unsafe.Pointer(&LongPtr_TotalNumberOfFreeBytes)),
		0, 0)

	if r != 0 {
		fmt.Printf(">>>> %s 的空间情况\n", diskName)
		fmt.Printf("剩余空间：%d G\n", LongPtr_FreeBytesAvailable/1024/1024/1024)
		fmt.Printf("用户可用空间：%d G\n", LongPtr_TotalNumberOfBytes/1024/1024/1024)
		fmt.Printf("剩余可用空间：%d G\n", LongPtr_TotalNumberOfFreeBytes/1024/1024/1024)

	} else {
		//此处的errno不是error接口，而是 type Errorno uintptr
		panic(errno)
	}
}

func main() {
	//释放方法引用
	defer syscall.FreeLibrary(kernel32)

	drives, err := getDriveNames()
	if err != nil {
		panic(err)
	}

	for _, d := range drives {
		//获取磁盘可用空间
		getDiskGreeSpace(d)
	}
}

/*syscall.Syscall系列方法
当前共5个方法

syscall.Syscall
syscall.Syscall6
syscall.Syscall9
syscall.Syscall12
syscall.Syscall15
分别对应 3个/6个/9个/12个/15个参数或以下的调用

参数都形如

syscall.Syscall(trap, nargs, a1, a2, a3)
第二个参数, nargs 即参数的个数,一旦传错, 轻则调用失败,重者直接APPCARSH

多余的参数, 用0代替*/

/*简单点的方式? 用syscall.Call
跟Syscall系列一样, Call方法最多15个参数. 这里用来Must开头的方法, 如不存在,会panic.

h := syscall.MustLoadDLL("kernel32.dll")
c := h.MustFindProc("GetDiskFreeSpaceExW")
lpFreeBytesAvailable := int64(0)
lpTotalNumberOfBytes := int64(0)
lpTotalNumberOfFreeBytes := int64(0)
r2, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("F:"))),
uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
uintptr(unsafe.Pointer(&lpTotalNumberOfBytes)),
uintptr(unsafe.Pointer(&lpTotalNumberOfFreeBytes)))
if r2 != 0 {
log.Println(r2, err, lpFreeBytesAvailable/1024/1024)
}*/
