package serial

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/sigurn/crc16"
	"io"
	"log"
	"strconv"
	"strings"
	//"test/src/mqtt"

	"encoding/gob"
)

type P struct {
	X, Y, Z int
	Name    string
}

type Q struct {
	X, Y *int32
	Name string
}

func main_____() {
	// Initialize the encoder and decoder.  Normally enc and dec would be
	// bound to network connections and the encoder and decoder would
	// run in different processes.
	var network bytes.Buffer        // Stand-in for a network connection
	enc := gob.NewEncoder(&network) // Will write to network.
	dec := gob.NewDecoder(&network) // Will read from network.
	// Encode (send) the value.
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	if err != nil {
		log.Fatal("encode error:", err)
	}

	// HERE ARE YOUR BYTES!!!!
	fmt.Println(network.Bytes())

	// Decode (receive) the value.
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error:", err)
	}
	fmt.Printf("%q: {%d,%d}\n", q.Name, *q.X, *q.Y)
}

func read(s io.ReadWriteCloser, uplink_Messages_t_up_topic string) {

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

			//	fmt.Printf("------------------ TEST %s", mqtt.Client)
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

}

func read__(date []byte) {

	r := bytes.NewReader(date)

	var len = make([]byte, 1)
	nr, errn := r.Read(len)
	if errn != nil {
		fmt.Println("binary.Read ==============%s", nr)
	}

	le := int(len[1])
	if le > 0 {
		return
	}
	var data__ struct {
		PI   float64
		Uate uint8
		Mine [8]byte
		Too  uint16
		Crc  uint16
	}

	var data__1 struct {
		Adr float64 //	Adr	1	读写器地址。

		reCmd uint8 //	reCmd	1	指示该响应数据块是哪个命令的应答。如果是对不可识别的命令的应答，则reCmd为0x00。

		Status [6]byte //	Status	1	命令执行结果状态值。

		Too uint16
		Crc uint16
	}

	fmt.Println(data__1.Adr)

	if err := binary.Read(r, binary.LittleEndian, &data__); err != nil {
		fmt.Println("binary.Read failed:", err)
	}
	fmt.Println("binary.Read ==============")

	fmt.Println(data__.PI)
	fmt.Println(data__.Uate)
	fmt.Printf("% x\n", data__.Mine)
	fmt.Println(data__.Too)
	fmt.Println("binary.Read ==============")

}
