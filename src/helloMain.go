package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"test/src/config"
	"test/src/execC"
	"test/src/mqtt"
	"test/src/update"
	"test/src/web"
	"text/template"

	//"test"

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
)

const letter = `
Dear {{.Name}},
{{if .Attended}}
It was a pleasure to see you at the wedding.
{{- else}}
It is a shame you couldn't make it to the wedding.
{{- end}}
{{with .Gift -}}
Thank you for the lovely {{.}}.
{{end}}
Best wishes,
Josie
`

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

var Input_pstrName = flag.String("name", "gerry", "input ur name")
var Input_piAge = flag.Int("age", 20, "input ur age")
var Input_flagvar int

func Init() {
	flag.IntVar(&Input_flagvar, "flagname", 1234, "help message for flagname")
}

func main() {

	// Prepare some data to insert into the template.
	type Recipient struct {
		Name, Gift string
		Attended   bool
	}
	var recipients = []Recipient{
		{"Aunt Mildred", "bone china tea set", true},
		{"Uncle John", "moleskin pants", false},
		{"Cousin Rodney", "", false},
	}

	// Create a new template and parse the letter into it.
	t := template.Must(template.New("letter").Parse(letter))

	// Execute the template for each recipient.
	for _, r := range recipients {
		err := t.Execute(os.Stdout, r)
		if err != nil {
			log.Println("executing template:", err)
		}
	}

	Init()
	flag.Parse()
	fmt.Println("name=", *Input_pstrName)
	fmt.Println("age=", *Input_piAge)
	fmt.Println("flagname=", Input_flagvar)

	if *Input_pstrName == "aaa" {

		var fileName = "e:\\docker-compose.yml"

		content, err := update.Install(fileName, "https://raw.githubusercontent.com/1026957152/test/master/src/docker-compose.yml")
		if err == nil {
			log.Printf("out, err := os.Create(filepath)" + content)

			templ := template.Must(template.New("compose").Parse(content))
			fmt.Println(templ.Name())

			// Create the file
			out, err := os.Create("e:\\docker-compose.yml__")
			if err != nil {
				panic(err)
			}
			log.Printf("out, err := os.Create(filepath)")
			defer out.Close()

			templ.Execute(os.Stdout, "Hello World")

			//templ, er :=template.ParseFiles(fileName)
			//	if er == nil {
			//	tt := template.Must(templ.Parse(letter))

			//	}
			execC.DockerCompose(fileName)
		}

	}

	/*	fyne.Fyncmain()
		return*/
	log.Printf("os.Create(filepath) docker")

	//update.DownloadFile_("e:\\a.txt","https://raw.githubusercontent.com/idreamsi/RadioHead/master/LICENSE")
	//	return

	macAddr := config.GetMacAddr()

	status["macAddr"] = macAddr
	status["deviceEui"] = macAddr + "ff"

	status["broker"] = "tcp://192.168.10.90:1883" // "tcp://mqtt.yulinmei.cn:1883"
	status["clientId"] = "storage_space_client_id" + macAddr + "ff"

	//    os.Exit(-1)
	/*   web.Webserver()

	     docker.CreateNewContainer("")



	    gpio.Start()*/
	_, cf := config.Open_config()
	//	imageName := "docker.yulinmei.cn/loan:0.0.1-SNAPSHOT"

	///docker.NewClient(cf.UpdateUrl) //"docker.io/library/alpine")

	go web.Webserver(&Wg, status)

	mqtt.New_mqtt(cf.AppID, status, cf.Server)
	//	mqtt.Mqtt_local(status)

	//	camera.ImageMain()

	//	r := strings.NewReplacer("<DevID>", status["deviceEui"], "<AppID>", cf.AppID)
	//qrcode.Qrcode_main(cf.Usbserial, r.Replace(mqtt.Uplink_Messages_t_up))

	//	return

	//serial.Seral_up_network(cfg)

	/*    var char string = "gb18030"
	      if mahonia.GetCharset(char) == nil {

	          fmt.Errorf("%s charset not suported \n", char)
	      }else{
	          fmt.Printf("%s charset  suported \n", char)

	      }*/
	Wg.Wait()

}
