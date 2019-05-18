package test

import (
	"bytes"
	"errors"
	"log"
	"net"
	"sync"
)

var ERR_EOF = errors.New("EOF")
var ERR_CLOSED_PIPE = errors.New("io: read/write on closed pipe")
var ERR_NO_PROGRESS = errors.New("multiple Read calls return no data or error")
var ERR_SHORT_BUFFER = errors.New("short buffer")
var ERR_SHORT_WRITE = errors.New("short write")
var ERR_UNEXPECTED_EOF = errors.New("unexpected EOF")

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

/*
func (self *AgentContext) CheckHostType(host_type string) error {
    switch host_type {
    case "virtual_machine":
        return nil
    case "bare_metal":
        return nil
    }
    return errors.New("CheckHostType ERROR:" + host_type)
}*/

var Wg sync.WaitGroup

func main() {
	var status = make(map[string]string) // map[string]string = {"macAddr",nil}

	macAddr := GetMacAddr()

	status["macAddr"] = macAddr
	status["deviceEui"] = macAddr + "ff"

	//    os.Exit(-1)
	/*   web.Webserver()

	     docker.CreateNewContainer("")



	    gpio.Start()*/
	cfg := Open_config()

	go Webserver(&Wg, status)

	log.Printf("MAIN 主程序继续 $v")

	//获取配置文件中的配置项
	_, err := cfg.String("TTN", "DeviceID")
	if err != nil {
		log.Printf("无法读取 $v", err)
	}

	server, err := cfg.String("TTN", "Server")
	if err != nil {
		log.Printf("无法读取 $v", err)
	}

	appID, err := cfg.String("TTN", "AppID")
	if err != nil {
		log.Printf("无法读取 $v", err)
	}

	New_mqtt(appID, status, server)
	log.Printf("MAIN 主程序继续  New_mqtt")

	Seral_up_network(cfg)
	log.Printf("MAIN 主程序继续 serial")

	NewClient()
	log.Printf("MAIN 主程序继续 docker")

	/*    var char string = "gb18030"
	      if mahonia.GetCharset(char) == nil {

	          fmt.Errorf("%s charset not suported \n", char)
	      }else{
	          fmt.Printf("%s charset  suported \n", char)

	      }*/
	Wg.Wait()
}
