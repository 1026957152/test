/*package gpio

import (
	"flag"
	"log"
	"os"
	"runtime"
)*/

package serial

import (
	"bytes"
	"fmt"
	//"github.com/karalabe/hid"
	//"flag"
	"github.com/larspensjo/config"
	"github.com/tarm/goserial"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

var (
	conFile = "\\src\\config.ini" //flag.String("configfile", "\\src\\config.ini", "config file")
)

var TOPIC = make(map[string]string)

func r() {

}

// getMacAddr gets the MAC hardware
// address of the host machine
func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			log.Printf("读取窗口信息%v  %s  %s", i.Flags, i.Name, i.HardwareAddr.String())

			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address

				addr = i.HardwareAddr.String()
				//break
			}
		}
	}
	return
}
func main() {
	//hid.Supported()

	addr := getMacAddr()
	log.Printf("读取窗口信息 %s", addr)

}
func main__() {
	//获取当前路径
	file, _ := os.Getwd()
	log.Printf("读取窗口信息 %s", file)
	log.Printf("读取窗口信息 %v", runtime.NumCPU())
	cfg, err := config.ReadDefault(file + conFile)

	if err != nil {
		log.Printf("无法找到", conFile, err)
	}
	//获取配置文件中的配置项
	id, err := cfg.String("COM", "COMID")
	if err != nil {
		log.Printf("无法读取 $v", err)
	}
	log.Printf("读取串口 %s", id)

	//设置串口编号
	c := &serial.Config{Name: id, Baud: 115200, ReadTimeout: time.Second * 2}

	//打开串口
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}

	// var commandn string
	var errn error
	var nw int
	//var buf byte[]
	//numbers := [2]int{1, 2}

	// CME ERROR: 10 SIM not inserted
	cmd := [18]string{
		//CME ERROR: 3:模块不支持该at指令。

		"AT+COPN\r",   //查询运营商名称,执行命令用于从ME返回运营商列表,
		"AT+CGREG?\r", //检测是否登陆上GPRS 网络
		"AT+CSQ\r",    //检测信号质量，确定是否可以登陆上网络,//99表示信道无效
		"at+ccid\r",   //检测是否装有SIM 卡
		"AT+CGSN\r",   // 获得GSM模块的IMEI（国际移动设备标识）序列号
		"AT+CGMR\r",   ////检测软件版本，5.0 以上的才有GPRS 功能支持
		"AT+CGMM\r",

		"AT\r",
		"AT+CIMI\r",
		"AT+CGDCONT=1,\"IP\"\r",
		/*				//AT+CGDCONT=1,"IP","CMNET"
						//1:表示使用第一种配置方案
						//IP:表示协议
						//CMNET:APN*/
		"AT+CGACT=1,1\r",   //激活网络 激活,返回OK则继续,  是激活PDP，建立modem和GPRS网络之间的连接
		"AT+CGCONTRDP=1\r", // 读取网络配置的IP地址/DNS，以及P-CSCF地址
		"AT+CGPADDR=1\r",   // 显示PDP地址

		"AT+CEREG=1",  //	EPS网络注册状态
		"AT+CEREG?\r", //检测网络注册状态

		"AT+CGREG?\r", //检测是否登陆上GPRS 网络

		"AT+CFUN=1\r",
		//AT+CFUN= 0,
		//modem不可以打电话，发短信，但是可以有其他操作，比如读 sim卡之类的。
		//AT+CFUN= 1,
		//modem可以打电话，发短信...所以叫做full functionality
		//比如打开射频可以用CFUN=1,关闭可以用CFUN=0,也可以自己定义打开蓝牙CFUN=5等等,所以具体需要看片子的说明说册,不是所有的人都完全按照3GPP规范来做的.
		//同理,Minimum functionality和Full functionality也是由各OEM自己定义的设备功能.
		"T+ZGACT=1,1\r", //                    若GEREG注册有效则能正常返回
	}

	aa := [14]string{
		"AT+CGDCONT=1,\"IP\"\r",
		"AT+CFUN=1\r",    //模块功能全打开，上电可以设置默认状态
		"AT+CEREG=1\r",   //注册上4G网络
		"AT+CGREG?\r",    //检测是否登陆上GPRS 网络
		"AT+CEREG?\r",    ////检测网络注册状态 查询３Ｇ使用
		"AT+ZGACT=1,1\r", // 若GEREG注册有效则能正常返回
		"AT+CGPADDR=1\r",
	}

	var count int = 0
	var step int = 7
	log.Printf(cmd[0] + " ")

	for _, x := range aa[count:step] {

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

}

func Seral_up_network(cfg *config.Config) {
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
	id, err := cfg.String("COM", "COMID")
	if err != nil {
		log.Printf("无法读取 $v", err)
	}
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

	// CME ERROR: 10 SIM not inserted
	cmd := [18]string{
		//CME ERROR: 3:模块不支持该at指令。

		"AT+COPN\r",   //查询运营商名称,执行命令用于从ME返回运营商列表,
		"AT+CGREG?\r", //检测是否登陆上GPRS 网络
		"AT+CSQ\r",    //检测信号质量，确定是否可以登陆上网络,//99表示信道无效
		"at+ccid\r",   //检测是否装有SIM 卡
		"AT+CGSN\r",   // 获得GSM模块的IMEI（国际移动设备标识）序列号
		"AT+CGMR\r",   ////检测软件版本，5.0 以上的才有GPRS 功能支持
		"AT+CGMM\r",

		"AT\r",
		"AT+CIMI\r",
		"AT+CGDCONT=1,\"IP\"\r",
		/*				//AT+CGDCONT=1,"IP","CMNET"
						//1:表示使用第一种配置方案
						//IP:表示协议
						//CMNET:APN*/
		"AT+CGACT=1,1\r",   //激活网络 激活,返回OK则继续,  是激活PDP，建立modem和GPRS网络之间的连接
		"AT+CGCONTRDP=1\r", // 读取网络配置的IP地址/DNS，以及P-CSCF地址
		"AT+CGPADDR=1\r",   // 显示PDP地址

		"AT+CEREG=1",  //	EPS网络注册状态
		"AT+CEREG?\r", //检测网络注册状态

		"AT+CGREG?\r", //检测是否登陆上GPRS 网络

		"AT+CFUN=1\r",
		//AT+CFUN= 0,
		//modem不可以打电话，发短信，但是可以有其他操作，比如读 sim卡之类的。
		//AT+CFUN= 1,
		//modem可以打电话，发短信...所以叫做full functionality
		//比如打开射频可以用CFUN=1,关闭可以用CFUN=0,也可以自己定义打开蓝牙CFUN=5等等,所以具体需要看片子的说明说册,不是所有的人都完全按照3GPP规范来做的.
		//同理,Minimum functionality和Full functionality也是由各OEM自己定义的设备功能.
		"T+ZGACT=1,1\r", //                    若GEREG注册有效则能正常返回
	}

	aa := [14]string{
		"AT+CGDCONT=1,\"IP\"\r",
		"AT+CFUN=1\r",    //模块功能全打开，上电可以设置默认状态
		"AT+CEREG=1\r",   //注册上4G网络
		"AT+CGREG?\r",    //检测是否登陆上GPRS 网络
		"AT+CEREG?\r",    ////检测网络注册状态 查询３Ｇ使用
		"AT+ZGACT=1,1\r", // 若GEREG注册有效则能正常返回
		"AT+CGPADDR=1\r",
	}

	var count int = 0
	var step int = 7
	log.Printf(cmd[0] + " ")

	for _, x := range aa[count:step] {

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

}
