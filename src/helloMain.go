package main

import (
	"bytes"
	"camera"
	"config"
	"docker"
	"errors"
	"mqtt"
	"qrcode"
	"strings"

	/*	"qrcode"
		"strings"*/

	/*	"fmt"
		"io"
		"net/http"
		"os"*/
	"log"
	//"mqtt"
	"net"

	"sync"
	"web"
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
var status = make(map[string]string) // map[string]string = {"macAddr",nil}

func main() {

	/*	fyne.Fyncmain()
		return*/
	log.Printf("os.Create(filepath) docker")

	//update.DownloadFile_("e:\\a.txt","https://raw.githubusercontent.com/idreamsi/RadioHead/master/LICENSE")
	//	return

	imageName := "docker.yulinmei.cn/loan:0.0.1-SNAPSHOT"

	docker.NewClient(imageName) //"docker.io/library/alpine")

	log.Printf("MAIN 主程序继续 docker")

	macAddr := config.GetMacAddr()

	status["macAddr"] = macAddr
	status["deviceEui"] = macAddr + "ff"

	status["broker"] = "tcp://192.168.10.90:1883" // "tcp://mqtt.yulinmei.cn:1883"
	status["clientId"] = "storage_space_client_id" + macAddr + "ff"

	//    os.Exit(-1)
	/*   web.Webserver()

	     docker.CreateNewContainer("")



	    gpio.Start()*/
	cfg := config.Open_config()

	go web.Webserver(&Wg, status)

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
	log.Printf("MAIN 主程序继续 docker", appID, status, server)
	mqtt.New_mqtt(appID, status, server)
	//	mqtt.Mqtt_local(status)

	camera.ImageMain()

	r := strings.NewReplacer("<DevID>", status["deviceEui"], "<AppID>", appID)
	qrcode.Qrcode_main(r.Replace(mqtt.Uplink_Messages_t_up))

	//	return

	log.Printf("MAIN 主程序继续  New_mqtt")

	//	serial.Seral_up_network(cfg)
	log.Printf("MAIN 主程序继续 serial")

	//	docker.NewClient("docker.io/library/alpine")

	log.Printf("MAIN 主程序继续 docker")

	/*    var char string = "gb18030"
	      if mahonia.GetCharset(char) == nil {

	          fmt.Errorf("%s charset not suported \n", char)
	      }else{
	          fmt.Printf("%s charset  suported \n", char)

	      }*/
	Wg.Wait()

}
