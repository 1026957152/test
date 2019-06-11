package serial

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/imroc/biu"
	"github.com/sigurn/crc16"
	"io"
	"unsafe"

	/*	"fmt"
		"io"*/

	//"github.com/karalabe/hid"
	//"flag"
	"github.com/tarm/goserial"
	"log"

	"time"
)

/*var (
	conFile = "\\src\\config.ini" //flag.String("configfile", "\\src\\config.ini", "config file")
)
*/
//piByte := []byte{0x18, 0x2d, 0x44, 0x54, 0xfb, 0x21, 0x09, 0x40}
//boolByte := []byte{0x00}

// baseBlockHdrEncoded is the wire encoded bytes of baseBlockHdr.
/*baseBlockHdrEncoded_ := []byte{
0x01, 0x00, 0x00, 0x00, // Version 1
0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, // PrevBlock
0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a, // MerkleRoot
0x29, 0xab, 0x5f, 0x49, // Timestamp
0xff, 0xff, 0x00, 0x1d, // Bits
0xf3, 0xe0, 0x01, 0x00, // Nonce
}

*/

/*inputs := [][]byte{
[]byte{0x81, 0x01},
[]byte{0x7f},
[]byte{0x03},
[]byte{0x01},
[]byte{0x00},
[]byte{0x02},
[]byte{0x04},
[]byte{0x7e},
[]byte{0x80, 0x01},
}
*/

type cmd struct {
	Len    []byte
	Adr    []byte
	Cmd    []byte
	Data   []byte
	CRC_16 []byte
}

// WAVHeader fixed-size datastructure
type WAVHeader struct {
	ChunkID       uint32
	ChunkSize     uint32
	Format        uint32
	Subchunk1ID   uint32
	Subchunk1Size uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   uint32
	Subchunk2Size uint32
}

type WAVFormat struct {
	Header WAVHeader
	Data   []byte
}

// Decode reads a WAVFormat from rd.
/*func Decode(rd io.Reader, dst *WAVFormat) error {
	if err := binary.Read(rd, binary.BigEndian, dst.Header); err != nil {
		return err
	}

	// Reuse the Data slice, if possible
	if cap(dst.Data) >= dst.Subchunk2Size {
		dst.Data = dst.Data[:dst.Subchunk2Size]
	} else {
		dst.Data = make([]byte, dst.Subchunk2Size)
	}

	if _, err := io.ReadFull(r, dst.Data); err != nil {
		return err
	}
	return nil
}*/

/*func r() {


	var cmd cmd
	cmd.Adr = []byte{0x00}
	cmd.Data = []byte{0x00}
	cmd.Cmd = []byte{0x10,0x12}
//	cmd.CRC_16 =
//	cmd.Len = Size(cmd)

}*/

type binData struct {
	A int32
	B int32
	C int16
}

/*func main() {
	fp, err := os.Open("tst.bin")

	if err != nil {
		panic(err)
	}

	defer fp.Close()
	for {
		thing := binData{}
		err := binary.Read(fp, binary.LittleEndian, &thing)
		if err == io.EOF{
			break
		}
		fmt.Println(thing.A, thing.B, thing.C)
	}
}
*/

type Header_without_crc struct {
	Length  uint8
	Address uint8
	Command uint8
}

type Header struct {
	Length  uint8
	Address uint8
	Command uint8
	/*	Data  [1]uint8
		CRC_16 uint16*/

}

type Reponse struct {
	Length uint8
	Adr    uint8
	ReCmd  uint8
	Status uint8
	/*	Data  [1]uint8
		CRC_16 uint16*/
}

type ReponseData读取读写器信息 struct {
	Version uint16 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Type    uint8  //Type	1	读写器类型代号。0x09代表UHFREADER18。
	Tr_Type uint8  //Tr_Type	1	读写器支持的协议信息，Bit1为1表示支持18000-6c协议， Bit0为1表示18000-6B协议，其它位保留。
	Dmaxfre uint8  //dmaxfre	1	Bit7-Bit6用于频段设置用；Bit5-Bit0表示当前读写器工作的最大频率。
	Dminfre uint8  //dminfre	1	Bit7-Bit6用于频段设置用；Bit5-Bit0表示当前读写器工作的最小频率。
	Power   uint8  //Power	1	读写器的输出功率。范围是0到30。
	Scntm   uint8  //Scntm	1	询查时间。读写器收到询查命令后，在询查时间内，会给上位机应答。
	/*	Data  [1]uint8
		CRC_16 uint16*/
}

type commandElement struct {
	C           uint8
	Description string
}

var command = map[string]commandElement{
	"寻查命令(单张)_18000-6B": {0x50, "寻查命令(单张)_18000-6B"},

	"询查标签_EPC_C1G2": {0x01, "询查标签_EPC_C1G2"},

	"读取读写器信息": {0x21, "读取读写器信息"},

	"设置串口波特率":  {0x28, "设置串口波特率"},
	"读取":       {0x21, "读取"},
	"调整功率":     {0x2F, "调整功率"},
	"声光控制命令":   {0x33, "声光控制命令"},
	"工作模式设置命令": {0x35, "工作模式设置命令"},
	"读取工作模式参数": {0x36, "读取工作模式参数"},
}

func Test(person string) (work func() string) {
	/*
	   Do someting
	*/
	work = func() string {
		return (person + " is working")
	}
	return
}

type FreqBand int32

const (
	User_band     FreqBand = 0
	Chinese_band2 FreqBand = 1
	US_band       FreqBand = 2
	Korean_band   FreqBand = 3
)

/*FreqBand {

Chinese_band2
US_band
Korean_band
}
MaxFre(Bit7)	MaxFre(Bit6)	MinFre(Bit7)	MinFre(Bit6)	FreqBand
0	0	0	0	User band
0	0	0	1	Chinese band2
0	0	1	0	US band
0	0	1	1	Korean band
0	1	0	0	保留
0	1	0	1	保留
…	…	…	…	…
1	1	1	1	保留
*/

func 设置读写器工作频率(freqBand FreqBand) (bb []byte, work func(client uint16, msg uint16)) {
	var (
		/**
		1个字节=8个二进制位,每种数据类型占用的字节数都不一样
		注意位操作千万不要越界了，如某个类型占8个bit位，偏移时候不要超过这个范围
		*/
		a uint8 = 30
	)
	//a输出结果:00011110
	fmt.Println(biu.ToBinaryString(a))
	/**
	将某一位设置为1，例如设置第8位，从右向左数需要偏移7位,注意不要越界
	1<<7=1000 0000 然后与a逻辑或|,偏移后的第8位为1，逻辑|运算时候只要1个为真就为真达到置1目的
	*/
	b := a | (1 << 7)
	//b输出结果:10011110
	fmt.Println(biu.ToBinaryString(b))
	/**
	将某一位设置为0，例如设置第4位，从右向左数需要偏移3位,注意不要越界
	1<<3=0000 1000 然后取反得到 1111 0111 然后逻辑&a
	*/
	c := a &^ (1 << 3)
	//c输出结果:00010110
	fmt.Println(biu.ToBinaryString(c))
	/**
	  获取某一位的值,即通过左右偏移来将将某位的值移动到第一位即可，当然也可以通过计算获得
	  如获取a的第4位
	  先拿掉4位以上的值 a<<4=1110 0000,然后拿掉右边的3位即可 a>>7=0000 0001
	*/
	d := (a << 4) >> 7
	//d输出结果:00000001 即1
	fmt.Println(biu.ToBinaryString(d))
	/**
	  取反某一位，即将某一位的1变0，0变1
	  这里使用到了亦或操作符 ^ 即 位值相同位0，不同为1
	  如获取a的第4位 1<<3=0000 1000
	  0000 1000 ^ 0001 1110 = 0001 0110
	*/
	e := a ^ (1 << 3)
	//d输出结果:00010110 即1
	fmt.Println(biu.ToBinaryString(e))

	//byte/[]byte -> string
	bs := []byte{1, 2, 3}
	s := biu.BytesToBinaryString(bs)
	fmt.Println(s)                               //[00000001 00000010 00000011]
	fmt.Println(biu.ByteToBinaryString(byte(3))) //00000011

	//string -> []byte
	//s := "[00000011 10000000]"
	//	bs := biu.BinaryStringToBytes(s)
	fmt.Printf("%#v\n", bs) //[]byte{0x3, 0x80}

	var MaxFre_Bit7 uint8 = 1
	var MaxFre_Bit6 uint8 = 2
	var MinFre_Bit7 uint8 = 3
	var MinFre_Bit6 uint8 = 4

	switch freqBand {
	case User_band:
		MaxFre_Bit7 = 1
		MaxFre_Bit6 = 1
		MinFre_Bit7 = 1
		MinFre_Bit6 = 2
		//return "MIN"
	case Chinese_band2:
		MaxFre_Bit7 = 1
		MaxFre_Bit6 = 1
		MinFre_Bit7 = 1
		MinFre_Bit6 = 2
	//	return "MAX"
	case US_band:
		MaxFre_Bit7 = 1
		MaxFre_Bit6 = 1
		MinFre_Bit7 = 1
		MinFre_Bit6 = 2
	//	return "MID"
	case Korean_band:
		MaxFre_Bit7 = 1
		MaxFre_Bit6 = 1
		MinFre_Bit7 = 1
		MinFre_Bit6 = 2
	//	return "AVG"
	default:
		MaxFre_Bit7 = 1
		MaxFre_Bit6 = 1
		MinFre_Bit7 = 1
		MinFre_Bit6 = 2
		//	return "UNKNOWN"
	}
	fmt.Println(biu.ToBinaryString(MaxFre_Bit7))
	fmt.Println(biu.ToBinaryString(MaxFre_Bit6))
	fmt.Println(biu.ToBinaryString(MinFre_Bit7))
	fmt.Println(biu.ToBinaryString(MinFre_Bit6))

	var f = func(client uint16, msg uint16) {

	}

	var cmd = command["调整功率"].C
	var dat = []uint8{0x10} //[]uint8{0x05}
	dat[0] = MinFre_Bit6

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	//b := make([]byte, 2)
	binary.LittleEndian.PutUint16(bb, checksum)
	bin_buf_without_crc.Write(bb)
	return bin_buf_without_crc.Bytes(), f
}

func 调整功率(Pwr uint8) (b []byte, work func(client uint16, msg uint16)) {

	var f = func(client uint16, msg uint16) {

	}

	var cmd = command["调整功率"].C
	var dat = []uint8{0x10} //[]uint8{0x05}
	dat[0] = Pwr

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	//b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f
}

func A声光控制命令(ActiveT uint8, SilentT uint8, Times uint8) (bb []byte, work func(client uint16, msg uint16), element commandElement) {

	var f = func(client uint16, msg uint16) {

	}

	element = command["声光控制命令"]
	var cmd = command["声光控制命令"].C
	var dat = []uint8{0x10, 0x10, 0x01} //[]uint8{0x05}
	dat[0] = ActiveT
	dat[1] = SilentT
	dat[2] = Times
	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element
}

func 读取工作模式参数() (bb []byte, work func(client uint16, msg uint16), cmd uint8) {
	var f = func(client uint16, msg uint16) {
	}

	cmd = command["读取工作模式参数"].C
	var dat = []uint8{} //[]uint8{0x05}

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, cmd

}

func 读取读写器信息() (bb []byte, work func(client uint16, msg uint16)) {
	var f = func(client uint16, msg uint16) {
	}

	var cmd = command["读取读写器信息"].C
	var dat = []uint8{} //[]uint8{0x05}

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f

}

func 寻查命令_单张_18000_6B() (bb []byte, work func(client uint16, msg uint16), element commandElement) {
	var f = func(client uint16, msg uint16) {
	}

	element = command["寻查命令(单张)_18000-6B"]
	var cmd = command["寻查命令(单张)_18000-6B"].C
	var dat = []uint8{} //[]uint8{0x05}

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func 询查标签_EPC_C1G2(AdrTID uint8, LenTID uint8) (bb []byte, work func(client uint16, msg uint16), element commandElement) {
	var f = func(client uint16, msg uint16) {
	}
	var dat = []uint8{} //[]uint8{0x05}
	if AdrTID == 0xFF && LenTID == 0xFF {

	} else {
		dat[0] = AdrTID
		dat[1] = LenTID
	}
	element = command["询查标签_EPC_C1G2"]
	var cmd = command["询查标签_EPC_C1G2"].C

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(dat)
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func main_______() {

	//设置串口编号
	c := &serial.Config{Name: "COM31", Baud: 57600, ReadTimeout: time.Second * 2}

	//打开串口
	s, err := serial.OpenPort(c)

	if err != nil {
		//log.Fatal(err)
		panic(err)

	}

	//bin_buf_without_crc,_,element := A声光控制命令(0x10,0x10,0x01 )
	//bin_buf_without_crc,_,_ := 读取工作模式参数()
	//bin_buf_without_crc,_ := 读取读写器信息()
	//bin_buf_without_crc,_,element := 寻查命令_单张_18000_6B()

	bin_buf_without_crc, _, element := 询查标签_EPC_C1G2(0xFF, 0xFF)

	time.Sleep(time.Second * 2)
	// 写入货柜串口命令

	nw, err := s.Write(bin_buf_without_crc)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("执行命令:%s %#[2]x", element.Description, element.C)
	log.Printf("成功写入 %d byte,发送数据:%x", nw, bin_buf_without_crc)

	time.Sleep(time.Second * 1)

	var nr int
	//var err_r error
	var buf = make([]byte, 1280)
	for i := 0; i < 1; i++ {
		log.Printf("开始读取")
		nr, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("成功读取  %d byte,读取内容:%x", nr, buf[:nr])

		checksum__ := crc16.Checksum(buf[:nr-2], crc16.MakeTable(crc16.CRC16_MCRF4XX))

		int16buf := new(bytes.Buffer)
		binary.Write(int16buf, binary.LittleEndian, checksum__)

		log.Printf("checksum__ %x , 校验的数据是 %x", int16buf, buf[nr-2:nr])

		if !bytes.Equal(int16buf.Bytes(), buf[nr-2:nr]) {

			return
		} else {
			log.Printf(" 校验的数据成功")

		}
		response, b := returnPrefix(buf, nr)

		if element.C == command["询查标签_EPC_C1G2"].C {

			if response.Status == 0x01 {
				log.Printf("成功读取  %x 个标签", b[:1])
				//log.Printf("成功读取  %d byte,读取内容:%x", nr, b[:nr])
				y := int(b[:1][0])

				initIndex := 1
				for i = 0; i < y; i++ {
					log.Printf("读取一个  %x 个标签", i)
					len := int(b[initIndex : initIndex+1][0])

					log.Printf("读取一个  %x 个标签", b[initIndex+1:initIndex+1+len])
					initIndex = initIndex + 1 + len
				}

			}

			return
		}

		if response.Status == 0 {

			//	log.Printf("buf____4: %x", b)

			if response.ReCmd == command["读取读写器信息"].C {

				var reponseData读取读写器信息 ReponseData读取读写器信息
				buf3 := bytes.NewReader(b)

				err = binary.Read(buf3, binary.LittleEndian, &reponseData读取读写器信息)
				if err == nil {
					log.Printf("reponseData读取读写器信息:\n%+v\n%x", reponseData读取读写器信息, reponseData读取读写器信息)
					log.Printf("reponseData读取读写器信息Dmaxfre:%b", reponseData读取读写器信息.Dmaxfre)
					log.Printf("reponseData读取读写器信息Dminfre:%b", reponseData读取读写器信息.Dminfre)
				} else {
					log.Printf("reponseData读取读写器信息:错误 %s", err)
				}

				log.Printf("读取内容 %x", buf[:nr])
			} else {
				log.Printf("其他事情，")

			}

			//log.Printf("读取内容 %s", buf[:nr])
		} else {

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		//var buf____4 = make([]byte, 128)
		//copy(buf____4,buf[unsafe.Sizeof(reponse):nr])

		//	log.Printf("%q", buf[:nr])
	}

	/*
		fmt.Printf("in = %#v\n", header)
		buf := new(bytes.Buffer)

		err := binary.Write(buf, binary.LittleEndian, header)
		if err != nil {
			log.Fatalf("binary.Write failed: %v", err)
		}
		b := buf.Bytes()
		fmt.Printf("wire = % x\n", b)

		var header2 Header
		buf2 := bytes.NewReader(b)
		err = binary.Read(buf2, binary.LittleEndian, &header2)
		if err != nil {
			log.Fatalf("binary.Read failed: %v", err)
		}
		fmt.Printf("out = %#v\n", header2)*/

}

var s io.ReadWriteCloser

func Init(epc_id []byte) (err error) {
	//设置串口编号
	c := &serial.Config{Name: "COM31", Baud: 57600, ReadTimeout: time.Second * 2}
	//打开串口
	s, err = serial.OpenPort(c)

	if err != nil {
		//log.Fatal(err)
		panic(err)

	}
	return err
}

func Send() (epc_id []byte, err error) {

	bin_buf_without_crc, _, element := 询查标签_EPC_C1G2(0xFF, 0xFF)

	time.Sleep(time.Second * 2)
	// 写入货柜串口命令

	nw, err := s.Write(bin_buf_without_crc)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("执行命令:%s %#[2]x", element.Description, element.C)
	log.Printf("成功写入 %d byte,发送数据:%x", nw, bin_buf_without_crc)

	time.Sleep(time.Second * 1)

	var nr int
	//var err_r error
	var buf = make([]byte, 1280)
	for i := 0; i < 1; i++ {
		log.Printf("开始读取")
		nr, err = s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("成功读取  %d byte,读取内容:%x", nr, buf[:nr])

		checksum__ := crc16.Checksum(buf[:nr-2], crc16.MakeTable(crc16.CRC16_MCRF4XX))

		int16buf := new(bytes.Buffer)
		binary.Write(int16buf, binary.LittleEndian, checksum__)

		log.Printf("checksum__ %x , 校验的数据是 %x", int16buf, buf[nr-2:nr])

		if !bytes.Equal(int16buf.Bytes(), buf[nr-2:nr]) {

			return
		} else {
			log.Printf(" 校验的数据成功")

		}
		response, b := returnPrefix(buf, nr)

		if element.C == command["询查标签_EPC_C1G2"].C {

			if response.Status == 0x01 {
				log.Printf("成功读取  %x 个标签", b[:1])
				//log.Printf("成功读取  %d byte,读取内容:%x", nr, b[:nr])
				y := int(b[:1][0])

				initIndex := 1
				for i = 0; i < y; i++ {
					log.Printf("读取一个  %x 个标签", i)
					len := int(b[initIndex : initIndex+1][0])

					log.Printf("读取一个  %x 个标签", b[initIndex+1:initIndex+1+len])
					epc_id = b[initIndex+1 : initIndex+1+len]

					initIndex = initIndex + 1 + len
				}

			}

			return epc_id, err
		}

	}

	/*
		fmt.Printf("in = %#v\n", header)
		buf := new(bytes.Buffer)

		err := binary.Write(buf, binary.LittleEndian, header)
		if err != nil {
			log.Fatalf("binary.Write failed: %v", err)
		}
		b := buf.Bytes()
		fmt.Printf("wire = % x\n", b)

		var header2 Header
		buf2 := bytes.NewReader(b)
		err = binary.Read(buf2, binary.LittleEndian, &header2)
		if err != nil {
			log.Fatalf("binary.Read failed: %v", err)
		}
		fmt.Printf("out = %#v\n", header2)*/

	return epc_id, err

}

var StatusMap = map[uint8]string{
	0xFE: "不合法的命令	当上位机输入的命令是不可识别的命令，如不存在的命令、或是CRC错误的命令",
	0xfb: "无电子标签可操作	当读写器对电子标签进行操作时，有效范围内没有可操作的电子标签时返回给上位机的状态值",
}

func returnPrefix(buf []byte, nr int) (reponse Reponse, b []byte) {
	buf2 := bytes.NewReader(buf)
	err := binary.Read(buf2, binary.LittleEndian, &reponse)
	if err == nil {
		//log.Printf("reponse: %x", reponse)

		log.Printf("返回内容分析: %x ：Length: %x，Adr: %x，ReCmd: %x，Status: %x", reponse, reponse.Length, reponse.Adr, reponse.ReCmd, reponse.Status)

	}
	//log.Printf("读取内容 reponse.Length: %d      nr:%d", unsafe.Sizeof(reponse),nr)
	//log.Printf("读取内容 reponse.Length: %x", buf[unsafe.Sizeof(reponse):nr])

	//	return bin_buf_without_crc.Bytes(),f

	return reponse, buf[unsafe.Sizeof(reponse):nr]

}
