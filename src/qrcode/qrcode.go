package qrcode

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/tarm/goserial"
	"io"
	"log"
	"strconv"
	"test/src/mqtt"

	//"os"
	"strings"

	//"github.com/stellar/go/crc16"
	"github.com/sigurn/crc16"
)

var TOPIC = make(map[string]string)

var c string = "$201003-881D"   //增加结束符CRLF 0x0D, 0x0A
var crc string = "$202301-9B71" //开启crc使用
func r(s io.ReadWriteCloser, uplink_Messages_t_up_topic string) {

	//var buffer = bytes.NewBuffer([]byte{0xad,0x12})

	//	var lenth = make([]byte, 1)
	//	nr, errn := s.Read(lenth)

	var data__ struct {
		PI   float64
		Uate uint8
		Mine [3]byte
		Too  uint16
	}
	//	r := bytes.NewReader(b)

	if err := binary.Read(s, binary.LittleEndian, &data__); err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("binary.Read ==============")

	fmt.Println(data__.PI)
	fmt.Println(data__.Uate)
	fmt.Printf("% x\n", data__.Mine)
	fmt.Println(data__.Too)
	fmt.Println("binary.Read ==============")

	//var nr int
	//	var errn error
	var buf = make([]byte, 1024)
	log.Printf("begin read")
	//	for i := 0; i < 1; i++ {
	//	nr, errn = s.Read(buf)
	//	log.Printf("读取结果 %q", buf[:nr])
	//	if errn != nil {
	//		log.Fatal(errn)
	//	}

	//	log.Printf("end read n:" + strconv.Itoa(nr))

	/*		if errn != nil {
				log.Fatal(errn)
			}
			log.Printf("end read n:" + strconv.Itoa(nr))
			log.Printf("读取窗口信息 %s", buf[:nr])
			log.Printf("读取结果 %q", buf[:nr])*/
	//}

	///data := []string{}
	scanner := bufio.NewScanner(s)

	/*
		split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
			//fmt.Printf("%t\t%d\t%s\n", atEOF, len(data), data)
			return 0, []byte{':'}, nil
		}
	*/
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		data := scanner.Text()
		//	data = append(data, scanner.Text())
		fmt.Printf("%s:\n", data)
		a := strings.Split(data, "TAIL")
		checksum := crc16.Checksum([]byte(a[0]+"TAIL"), crc16.MakeTable(crc16.CRC16_XMODEM))

		int16buf := new(bytes.Buffer)

		binary.Write(int16buf, binary.BigEndian, checksum)

		fmt.Printf("write buf is: %+X \n", int16buf.Bytes())

		fmt.Printf("%d****%s****%s****checksum%x\n", len(a), a[0], a[1], checksum)
		fmt.Printf("%s********checksum%x\n", a[0]+"TAIL", int16buf)

		//	fmt.Printf("%s********checksum%X\n", a[0], checksum_WITHOU_TAIL)
		//	fmt.Printf("%s********checksum%X\n", data, checksum_ALL)

		if strings.Contains(data, `"sensor"`) {
			//sensor <- data
		} else {
			//actuator <- data
		}

		var s = a[1]
		var base = 16
		var size = 16
		value, _ := strconv.ParseUint(s, base, size)

		if checksum == uint16(value) {
			var uplink_message map[string]interface{} = make(map[string]interface{})
			//	uplink_message["uplink_message"] = uplink_message
			uplink_message["session_key_id"] = "AWiZpAyXrAfEkUNkBljRoA=="
			uplink_message["uplink_token"] = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"
			uplink_message["pay_load"] = data

			fmt.Printf("------------------ TEST %s", mqtt.Client)
			mqtt.Uplink_Messages_t_up_mqtt(uplink_Messages_t_up_topic, uplink_message)
		}

	}

	for {
		fmt.Println("开始读：")

		n, err := s.Read(buf) // 这里读，持续读 10秒， 然后超时

		fmt.Println("等待读：")

		if err != nil {
			log.Fatal(err)
		}
		if n != 0 {
			s := string(buf[:n])
			fmt.Println(":" + s)

		}

	}

	/*	if(s.available() == 2){
		char command = Serial.read();
		int value = Serial.read();
		if(command == ‘s’) {
		myservo.write(value);
		}
		if(command == ‘l’) {
		// contro lights
		}
	}*/

}

func Qrcode_main(uplink_Messages_t_up_topic string) {
	fmt.Printf("打开串口：%s \n")

	//	file, _ := os.Getwd()

	//设置串口编号
	c := &serial.Config{Name: "COM30", Baud: 9600, ReadTimeout: 0} //time.Second * 1}
	//打开串口
	s, err := serial.OpenPort(c)

	if err != nil {
		log.Fatal(err)
	}
	//     var nr1 int
	//var err_r error
	//   var buf1 = make([]byte, 12800)

	go r(s, uplink_Messages_t_up_topic)

	//	input := bufio.NewScanner(os.Stdin)

	/*	for input.Scan() {
		// counts[input.Text()]++
		fmt.Printf("输入数据：%s \n",input.Text())


	}*/
	//log.Printf("写入  %d", nw)

	//                       nr1, _ = s.Read(buf1)
	///                      log.Printf("end read n:" + strconv.Itoa(nr1))

	//        if errn != nil {
	//      log.Fatal(errn)
	// }
	//                   log.Printf("end read n:" + strconv.Itoa(nr1))
	//                 log.Printf("读取窗口信息 %s", buf1[:nr1])
	//               log.Printf("读取结果 %q", buf1[:nr1])

}
