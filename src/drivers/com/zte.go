package com

import (
	"fmt"
	serial "github.com/tarm/goserial"
	"log"
	"os"
	"runtime"
	"strconv"
	"test/src/hikvision"
	"time"
)

// LedDriver represents a digital Led
type ZteDriver struct {
	pin        string
	name       string
	connection SerialWriter
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
func NewZteDriver(a SerialWriter, pin string) *ZteDriver {
	l := &ZteDriver{
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
func (l *ZteDriver) Start() (err error) { return }

// Halt implements the Driver interface
func (l *ZteDriver) Halt() (err error) { return }

// Name returns the LedDrivers name
func (l *ZteDriver) Name() string { return l.name }

// SetName sets the LedDrivers name
func (l *ZteDriver) SetName(n string) { l.name = n }

// Pin returns the LedDrivers name
func (l *ZteDriver) Pin() string { return l.pin }

// Connection returns the LedDrivers Connection
func (l *ZteDriver) Connection() hikvision.Connection {
	return l.connection.(hikvision.Connection)
}

// State return true if the led is On and false if the led is Off
func (l *ZteDriver) State() bool {
	return l.high
}

var Aa = [14]string{
	"AT+CGDCONT=1,\"IP\"\r",
	"AT+CFUN=1\r",    //模块功能全打开，上电可以设置默认状态
	"AT+CEREG=1\r",   //注册上4G网络
	"AT+CGREG?\r",    //检测是否登陆上GPRS 网络
	"AT+CEREG?\r",    ////检测网络注册状态 查询３Ｇ使用
	"AT+ZGACT=1,1\r", // 若GEREG注册有效则能正常返回
	"AT+CGPADDR=1\r",
}

// On sets the led to a high state.
func (l *ZteDriver) On() (err error) {
	// var commandn string
	var errn error
	var nw int
	//var buf byte[]
	//numbers := [2]int{1, 2}

	var count int = 0
	var step int = 7

	for _, x := range Aa[count:step] {

		time.Sleep(time.Second * 2)

		//	commandx := "COMMAND" + x //strconv.Itoa(x)
		//	_, _ = cfg.String("COM", commandx)
		// 写入货柜串口命令
		log.Printf("写入串口命令"+"  %s", x)

		if err = l.connection.SerialWriter([]byte(x)); err != nil {
			return
		}
		//l.connection.SerialWriter([]byte(x))

		nw, errn = s.Write([]byte(x))

		if errn != nil {
			log.Fatal(errn)
		}
		log.Printf("写入  %d", nw)

		time.Sleep(time.Second * 1)

		var nr int
		//var err_r error
		var buf = make([]byte, 128)
		for i := 0; i < 1; i++ {
			log.Printf("开始读取")

			nr, _ = s.Read(buf)
			log.Printf("结束读取" + strconv.Itoa(nr))

			// if errn != nil {
			//      log.Fatal(errn)
			// }

			log.Printf("读取内容 %s", buf[:nr])

			//	log.Printf("%q", buf[:nr])
		}
	}

	if err = l.connection.DigitalWrite(l.Pin(), 1); err != nil {
		return
	}
	l.high = true
	return
}

// Off sets the led to a low state.
func (l *ZteDriver) Off() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 0); err != nil {
		return
	}
	l.high = false
	return
}

// Toggle sets the led to the opposite of it's current state
func (l *ZteDriver) Toggle() (err error) {
	if l.State() {
		err = l.Off()
	} else {
		err = l.On()
	}
	return
}

// Brightness sets the led to the specified level of brightness
func (l *ZteDriver) Brightness(level byte) (err error) {
	if writer, ok := l.connection.(PwmWriter); ok {
		return writer.PwmWrite(l.Pin(), level)
	}
	return ErrPwmWriteUnsupported
}

func Seral_up_network(id string) (name string, err error) {
	log.Printf("MAIN 主程序继续 serial")

	defer func() {
		fmt.Println("Mqqt 找值，defer end...")
	}()
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()

	//获取当前路径
	file, _ := os.Getwd()
	log.Printf("读取窗口信息 %s", file)
	log.Printf("读取窗口信息 %v", runtime.NumCPU())

	//获取配置文件中的配置项
	/*	id, err := cfg.String("COM", "COMID")
		if err != nil {
			log.Printf("无法读取 $v", err)
		}*/
	log.Printf("读取串口 %s", id)

	//设置串口编号
	c := &serial.Config{Name: id, Baud: 115200, ReadTimeout: time.Second * 2}

	//打开串口
	s, err := serial.OpenPort(c)

	if err != nil {
		//log.Fatal(err)
		panic(err)

	}

	// var commandn string
	var errn error
	var nw int
	//var buf byte[]
	//numbers := [2]int{1, 2}

	var count int = 0
	var step int = 7

	for _, x := range Aa[count:step] {

		time.Sleep(time.Second * 2)

		//	commandx := "COMMAND" + x //strconv.Itoa(x)
		//	_, _ = cfg.String("COM", commandx)
		// 写入货柜串口命令
		log.Printf("写入串口命令"+"  %s", x)
		nw, errn = s.Write([]byte(x))

		if errn != nil {
			log.Fatal(errn)
		}
		log.Printf("写入  %d", nw)

		time.Sleep(time.Second * 1)

		var nr int
		//var err_r error
		var buf = make([]byte, 128)
		for i := 0; i < 1; i++ {
			log.Printf("开始读取")

			nr, _ = s.Read(buf)
			log.Printf("结束读取" + strconv.Itoa(nr))

			// if errn != nil {
			//      log.Fatal(errn)
			// }

			log.Printf("读取内容 %s", buf[:nr])

			//	log.Printf("%q", buf[:nr])
		}
	}

	return "4g", nil
}
