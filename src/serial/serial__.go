package serial

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/imroc/biu"
	"github.com/sigurn/crc16"
	"io"
	"os"
	"unsafe"

	/*	"fmt"
		"io"*/

	//"github.com/karalabe/hid"
	//"flag"
	"github.com/tarm/goserial"
	"log"
	"time"
)

type PolicyType int32

const (
	Policy_保留区    PolicyType = 0
	Policy_EPC存储区 PolicyType = 1
	Policy_TID存储区 PolicyType = 2
	Policy_用户存储区  PolicyType = 3
)

func (p PolicyType) String() uint8 {
	switch p {
	case Policy_保留区:
		return 0x00
	case Policy_EPC存储区:
		return 0x01
	case Policy_TID存储区:
		return 0x02
	case Policy_用户存储区:
		return 0x03
	default:
		return 0xff
	}
}

//选择要读取的存储区。0x00：保留区；0x01：EPC存储区；0x02：TID存储区；0x03：用户存储区。
func foo(p PolicyType) {
	fmt.Printf("enum value: %v\n", p)
}

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

type Data读数据_data struct {
	//	ENum uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	//EPC [12]byte //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Mem     uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	WordPtr uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号

	Num     uint8   //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Pwd     [4]byte //Type	1	读写器类型代号。0x09代表UHFREADER18。
	MaskAdr uint8   //Tr_Type	1	读写器支持的协议信息，Bit1为1表示支持18000-6c协议， Bit0为1表示18000-6B协议，其它位保留。
	MaskLen uint8   //dmaxfre	1	Bit7-Bit6用于频段设置用；Bit5-Bit0表示当前读写器工作的最大频率。
	/*	Data  [1]uint8
		CRC_16 uint16*/
}

type Data写数据_data struct {
	WNum    uint8
	ENum    uint8  //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	EPC     []byte //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Mem     uint8  //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	WordPtr uint8  //Version	2	版本号，高字节代表主版本号，低字节代表子版本号

	Wdt     uint8  //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Pwd     uint32 //Type	1	读写器类型代号。0x09代表UHFREADER18。
	MaskAdr uint8  //Tr_Type	1	读写器支持的协议信息，Bit1为1表示支持18000-6c协议， Bit0为1表示18000-6B协议，其它位保留。
	MaskLen uint8  //dmaxfre	1	Bit7-Bit6用于频段设置用；Bit5-Bit0表示当前读写器工作的最大频率。
	/*	Data  [1]uint8
		CRC_16 uint16*/
}

type Data工作模式设置命令_data struct {
	Read_mode  uint8
	Mode_state uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Mem_Inven  uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	First_Adr  uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号

	Word_Num uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Tag_Time uint8 //Type	1	读写器类型代号。0x09代表UHFREADER18。

}

type commandElement struct {
	C           uint8
	Description string
}

var command = map[string]commandElement{

	"读数据_18000-6B": {0x52, "读数据_18000-6B"},

	"寻查命令(单张)_18000-6B": {0x50, "寻查命令(单张)_18000-6B"},

	"询查多标签_EPC_C1G2": {0x01, "询查多标签_EPC_C1G2"},

	"读取读写器信息": {0x21, "读取读写器信息"},

	"设置串口波特率":                   {0x28, "设置串口波特率"},
	"读取":                        {0x21, "读取"},
	"调整功率":                      {0x2f, "调整功率"},
	"声光控制命令":                    {0x33, "声光控制命令"},
	"工作模式设置命令":                  {0x35, "工作模式设置命令"},
	"读取工作模式参数":                  {0x36, "读取工作模式参数"},
	"读数据_EPC_C1_G2_ISO18000-6C": {0x02, "读数据_EPC_C1_G2_ISO18000-6C"},

	"写EPC号_EPC_C1_G2_ISO18000-6C":  {0x04, "写EPC号_EPC_C1_G2_ISO18000-6C"},
	"询查单张标签_EPC_C1_G2_ISO18000-6C": {0x0f, "询查单张标签_EPC_C1_G2_ISO18000-6C"},

	"EAS检测精度设置":       {0x37, "EAS检测精度设置"},
	"设定存储区读写保护状态":     {0x06, "设定存储区读写保护状态"},
	"测试标签是否被设置读保护":    {0x0b, "测试标签是否被设置读保护"},
	"读保护设置(根据EPC号设定)": {0x08, "读保护设置(根据EPC号设定)"},
	"解锁读保护":           {0x0a, "解锁读保护"},
	"写数据":             {0x03, "写数据"},
	"块擦除":             {0x07, "块擦除"},

	"销毁标签": {0x05, "销毁标签"},
}

type Data设定存储区读写保护状态_data struct {

	//	ENum uint8
	//	EPC []byte //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	Select     uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	SetProtect uint8 //Version	2	版本号，高字节代表主版本号，低字节代表子版本号

	Pwd     [4]byte //Version	2	版本号，高字节代表主版本号，低字节代表子版本号
	MaskAdr uint8   //Type	1	读写器类型代号。0x09代表UHFREADER18。
	MaskLen uint8
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

func (p FreqBand) String() uint8 {
	switch p {
	case User_band:
		return 0x00
	case Chinese_band2:
		return 0x01
	case US_band:
		return 0x02
	case Korean_band:
		return 0x03
	default:
		return 0xff
	}
}

//选择要读取的存储区。0x00：保留区；0x01：EPC存储区；0x02：TID存储区；0x03：用户存储区。
func foo_FreqBand(p FreqBand) {
	fmt.Printf("enum value: %v\n", p)
}

type Read_mode_工作模式 int32

const (
	应答模式       Read_mode_工作模式 = 0
	主动模式       Read_mode_工作模式 = 1
	触发模式_低电平有效 Read_mode_工作模式 = 2
	触发模式_高电平有效 Read_mode_工作模式 = 3
)

/*0	0	应答模式
0	1	主动模式
1	0	触发模式(低电平有效)
1	1	触发模式(高电平有效)*/
func (p Read_mode_工作模式) String() uint8 {
	switch p {
	case 应答模式:

		return biu.BinaryStringToBytes("00000000")[0]
	case 主动模式:

		return biu.BinaryStringToBytes("00000001")[0]
	case 触发模式_低电平有效:

		return biu.BinaryStringToBytes("00000010")[0]
	case 触发模式_高电平有效:

		return biu.BinaryStringToBytes("00000011")[0]
	default:
		return 0xff
	}
}

type Select_mode int32

const (
	控制Kill密码读写保护设定 Select_mode = 0
	控制访问密码读写保护设定   Select_mode = 1
	控制EPC存储区读写保护设定 Select_mode = 2
	控制TID存储区读写保护设定 Select_mode = 3
	控制用户存储区读写保护设定  Select_mode = 4
)

/*Select为0x00时，控制Kill密码读写保护设定。
Select为0x01时，控制访问密码读写保护设定。
Select为0x02时，控制EPC存储区读写保护设定。
Select为0x03时，控制TID存储区读写保护设定。
Select为0x04时，控制用户存储区读写保护设定。*/
func (p Select_mode) String() uint8 {
	switch p {
	case 控制Kill密码读写保护设定:
		return 0x00
	case 控制访问密码读写保护设定:

		return 0x01
	case 控制EPC存储区读写保护设定:

		return 0x02
	case 控制TID存储区读写保护设定:
		return 0x03
	case 控制用户存储区读写保护设定:
		return 0x04

	default:
		return 0xff
	}
}

type Select_SetProtect_mode int32

const (
	Kill密码区或访问密码区_设置为无保护下的可读可写 Select_SetProtect_mode = 0
	Kill密码区或访问密码区_设置为永远可读可写    Select_SetProtect_mode = 1
	Kill密码区或访问密码区_设置为带密码可读可写   Select_SetProtect_mode = 2
	Kill密码区或访问密码区_设置为永远不可读不可写  Select_SetProtect_mode = 3

	设置EPC区_TID区及用户区_设置为无保护下的可写 Select_SetProtect_mode = 5
	设置EPC区_TID区及用户区_设置为永远可写    Select_SetProtect_mode = 6
	设置EPC区_TID区及用户区_设置为带密码可写   Select_SetProtect_mode = 7
	设置EPC区_TID区及用户区_设置为永远不可写   Select_SetProtect_mode = 8
)

/*
SetProtect：SetProtect的值根据Select的值而确定。
当Select为0x00或0x01，即当设置Kill密码区或访问密码区的时候，SetProtect的值代表的意义如下：
0x00：设置为无保护下的可读可写
0x01：设置为永远可读可写
0x02：设置为带密码可读可写
0x03：设置为永远不可读不可写

当Select为0x02、0x03、0x04的时候，即当设置EPC区、TID区及用户区的时候，SetProtect的值代表的意义如下：
0x00：设置为无保护下的可写
0x01：设置为永远可写
0x02：设置为带密码可写
0x03：设置为永远不可写
*/
func (p Select_SetProtect_mode) String() uint8 {
	switch p {
	case Kill密码区或访问密码区_设置为无保护下的可读可写:
		return 0x00
	case Kill密码区或访问密码区_设置为永远可读可写:
		return 0x01
	case Kill密码区或访问密码区_设置为带密码可读可写:
		return 0x02
	case Kill密码区或访问密码区_设置为永远不可读不可写:
		return 0x03

	case 设置EPC区_TID区及用户区_设置为无保护下的可写:
		return 0x00
	case 设置EPC区_TID区及用户区_设置为永远可写:
		return 0x01
	case 设置EPC区_TID区及用户区_设置为带密码可写:
		return 0x02
	case 设置EPC区_TID区及用户区_设置为永远不可写:
		return 0x03
	default:
		return 0xff

	}
}

/*Mode_state：Bit0：协议选择位。Bit0=0时读写器支持18000-6C协议；Bit0=1时读写器支持18000-6B协议。
Bit1：输出方式选择位。Bit1=0时韦根输出，Bit1=1时RS232/RS485输出。
Bit2：蜂鸣器提示选择位。Bit2=0时开蜂鸣器提示，Bit2=1时关蜂鸣器提示，默认值为0。
Bit3：韦根输出模式下First_Adr参数为字地址或字节地址选择位。Bit3=0
时First_Adr为字地址；Bit3=1时First_Adr为字节地址。
Bit4：玺瑞485选择位，Bit1=0时该位无效。Bit4=0时是普通485输出方式，Bit4=1时是玺瑞485模式。玺瑞485模式下只支持单标签操作（18000-6C、18000-6B均有效）（读保留 区、EPC区、TID区、用户区，单张查询）。玺瑞485模式下First_Adr为字节地址。
其它位保留，默认为0。*/

type Mode_state_工作模式 int32

const (
	协议选择位                          Mode_state_工作模式 = 0
	输出方式选择位                        Mode_state_工作模式 = 1
	蜂鸣器提示选择位                       Mode_state_工作模式 = 2
	韦根输出模式下First_Adr参数为字地址或字节地址选择位 Mode_state_工作模式 = 3
	玺瑞485选择位                       Mode_state_工作模式 = 4
)

/*0	0	应答模式
0	1	主动模式
1	0	触发模式(低电平有效)
1	1	触发模式(高电平有效)*/
func (p Mode_state_工作模式) String() uint8 {
	switch p {
	case 协议选择位:

		return biu.BinaryStringToBytes("00000000")[0]
	case 输出方式选择位:

		return biu.BinaryStringToBytes("00000001")[0]
	case 蜂鸣器提示选择位:

		return biu.BinaryStringToBytes("00000010")[0]
	case 韦根输出模式下First_Adr参数为字地址或字节地址选择位:

		return biu.BinaryStringToBytes("00000011")[0]
	case 玺瑞485选择位:

		return biu.BinaryStringToBytes("00000011")[0]
	default:
		return 0xff
	}
}

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

func __调整功率(Pwr uint8) (b__ []byte, fun func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		log.Printf("读取的数据 %x", b)

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {
			log.Printf("__调整功率 成功了")
			return nil, err

		} else {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		return nil, err

	}

	element = command["调整功率"]
	var cmd = element.C
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
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)

	return bin_buf_without_crc.Bytes(), f, element

}

func __A声光控制命令(ActiveT uint8, SilentT uint8, Times uint8) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		/*
			if element.C == command["询查多标签_EPC_C1G2"].C {
		*/
		if response.Status == 0x00 {
			log.Printf("A声光控制命令，读取  %x 个标签", b[:1])
			return epc_ids, err
		}

		return nil, err

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

func 读取工作模式参数() (bb []byte, work func(client uint16, msg uint16), element commandElement) {
	var f = func(client uint16, msg uint16) {
	}

	element = command["读取工作模式参数"]
	var dat = []uint8{} //[]uint8{0x05}

	var length = uint8(4 + len(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = element.C
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

func __读取读写器信息() (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			//	log.Printf("buf____4: %x", b)

			if response.ReCmd == element.C {

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

				//	log.Printf("读取内容 %x", buf[:nr])
			} else {
				log.Printf("其他事情，")

			}

			//log.Printf("读取内容 %s", buf[:nr])
		} else {

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err

	}

	element = command["读取读写器信息"]
	var cmd = element.C
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

func __询查多标签_EPC_C1G2(AdrTID uint8, LenTID uint8) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {
	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		/*
			if element.C == command["询查多标签_EPC_C1G2"].C {
		*/
		if response.Status == 0x01 {
			log.Printf("成功读取  %x 个标签", b[:1])
			//log.Printf("成功读取  %d byte,读取内容:%x", nr, b[:nr])

			buff := bytes.NewBuffer(b)

			//re := bytes.NewReader(b);
			buf_count := make([]byte, 1)
			_, _ = buff.Read(buf_count)

			count := int(buf_count[0])

			//返回未读取部分的长度
			//	fmt.Println("re len : ", re.Len());
			//返回底层数据总长度
			//	fmt.Println("re size : ", re.Size());

			epc_ids := make([][]byte, count)
			for i := 0; i < count; i++ {
				buf_epc_len := make([]byte, 1)
				buff.Read(buf_epc_len)
				len_epc := int(buf_epc_len[0])

				buf_epc_code := make([]byte, len_epc)
				buff.Read(buf_epc_code)

				log.Printf("读取一个  %x 个标签", i)

				log.Printf("读取一个 长度：%d 的标签 %x ", len_epc, buf_epc_code)
				epc_ids[i] = buf_epc_code

				//	fmt.Printf("----------- epc_ids_ %x\n", epc_ids[i])

				//	data_byte, _,element := 读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i])
				//	Send_(data_byte,element)
				//	fmt.Printf("----------- epc_ids_ %x", data_byte)

				/*				data_byte, _,element := ______写EPC号_EPC_C1_G2_ISO18000([4]byte{0x01,0x01,0x01,0x01},[]byte{0x01,0x01,0x01,0x01})
								Send_(data_byte,element)*/

				data_byte, fun, element := __读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i], Policy_用户存储区, 0x00, 0x08)
				Send_(data_byte, element, fun)
				fmt.Printf("----------- epc_ids_ %x", data_byte)
			}

			return epc_ids, err

		}

		return nil, err

	}

	var dat = []uint8{} //[]uint8{0x05}
	if AdrTID == 0xFF && LenTID == 0xFF {

	} else {
		dat[0] = AdrTID
		dat[1] = LenTID
	}
	element = command["询查多标签_EPC_C1G2"]
	var cmd = element.C

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

func __写EPC号_EPC_C1_G2_ISO18000(Pwd [4]byte, WEPC []byte) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		log.Printf("读取的数据 %x", b)

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			//	log.Printf("buf____4: %x", b)

			if response.ReCmd == element.C {

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

				//	log.Printf("读取内容 %x", buf[:nr])
			} else {
				log.Printf("其他事情，")

			}

			//log.Printf("读取内容 %s", buf[:nr])
		} else {

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err

	}

	/*	var dat = []uint8{} //[]uint8{0x05}
		dat[0] = ENum
	*/
	element = command["写EPC号_EPC_C1_G2_ISO18000-6C"]
	var cmd = element.C

	var length = uint8(4 + 1 + len(Pwd) + len(WEPC))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)

	/*		pwd := make([]byte, 4)
	binary.LittleEndian.PutUint32(pwd, Pwd)*/

	var ENum uint8 = uint8(len(WEPC) / 2)

	eNum := make([]byte, 1)
	eNum[0] = ENum
	bin_buf_without_crc.Write(eNum)
	bin_buf_without_crc.Write(Pwd[:])
	bin_buf_without_crc.Write(WEPC)

	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func __读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(EPC []byte, Mem PolicyType, WordPtr uint8, Num uint8) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	//	header_without_crc := &Header_without_crc{Address: 0x00}
	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD || response.Status == 0xFc {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0x00 {

			buff := bytes.NewBuffer(b[:len(b)-2])

			if Mem == Policy_EPC存储区 {
				log.Printf("Policy_EPC存储区  %x", buff.Bytes())
				log.Printf("Policy_EPC存储区  CRC-16:%x", buff.Bytes()[:2])
				log.Printf("Policy_EPC存储区 PC:%#b", buff.Bytes()[2:4])
				var bit = uint8(buff.Bytes()[2:3][0])
				bit = bit >> 3
				log.Printf("Policy_EPC存储区 10h—14h:%#b", bit)
				log.Printf("Policy_EPC存储区 epc长度:%d", bit)

				log.Printf("Policy_EPC存储区 EPC:%x", buff.Bytes()[4:])
			}
			if Mem == Policy_TID存储区 {
				log.Printf("Policy_TID存储区 成功读取串口返回数据  %x", buff.Bytes())

			}
			if Mem == Policy_保留区 {
				log.Printf("Policy_保留区 成功读取串口返回数据  %x", buff.Bytes())
				log.Printf("Policy_保留区  销毁密码，%x", buff.Bytes()[:4])
				log.Printf("Policy_保留区  访问密码%x", buff.Bytes()[5:])
			}
			if Mem == Policy_用户存储区 {
				log.Printf("Policy_用户存储区 成功读取串口返回数据  %x", buff.Bytes())
				log.Printf("Policy_用户存储区 成功读取串口返回数据  %x", buff.Bytes())
				log.Printf("Policy_用户存储区  销毁密码，%x", buff.Bytes()[:4])
				log.Printf("Policy_用户存储区  访问密码%x", buff.Bytes()[5:])
			}
			return [][]byte{[]byte{0x01}}, err

		}

		return [][]byte{[]byte{0x01}}, err

	}

	//data.ENum =   uint8(len(EPC)/2)
	var bin_buf_Data bytes.Buffer

	binary.Write(&bin_buf_Data, binary.BigEndian, uint8(len(EPC)/2))
	binary.Write(&bin_buf_Data, binary.BigEndian, EPC[:len(EPC)])

	/*	var arr  = make([]byte,len(EPC))
		copy(arr[:], EPC[:len(EPC)])
		data.EPC = arr*/

	/*	data.Mem= PolicyType.String(Mem)
		data.WordPtr = 0x00
	/*	if Mem == Policy_保留区 {
			data.Num = 0x04
		}else{
			data.Num = Num
		}
		data.Num = Num

		data.Pwd =
		data.MaskAdr = 0x00
		data.MaskLen = 0x01
	*/

	//var ENum uint8 = uint8(len(EPC)/2)
	mem := make([]byte, 1)
	mem[0] = PolicyType.String(Mem)
	bin_buf_Data.Write(mem)

	wordPtr := make([]byte, 1)
	wordPtr[0] = WordPtr
	bin_buf_Data.Write(wordPtr)

	num := make([]byte, 1)
	num[0] = Num
	bin_buf_Data.Write(num)
	bin_buf_Data.Write([]byte{0x00, 0x00, 0x00, 0x00})
	/*	bin_buf_Data.Write([]byte{0x00})
		bin_buf_Data.Write([]byte{0x00})*/

	//	data.MaskAdr = uint8(0)
	//	data.MaskLen = uint8(0)

	log.Printf("生成命令: %s", element.Description)
	//binary.Write(&bin_buf_data, binary.BigEndian, data)

	log.Printf("data[] 数据 :%x", bin_buf_Data.Bytes())
	log.Printf("data[] 数据 :%+x", bin_buf_Data.Bytes())
	log.Printf("data[] 数据 :%+v", bin_buf_Data.Bytes())

	/*	if AdrTID == 0xFF && LenTID == 0xFF {

		} else {
			dat[0] = AdrTID
			dat[1] = LenTID
		}*/
	element = command["读数据_EPC_C1_G2_ISO18000-6C"]
	var cmd = element.C

	var length = uint8(4 + len(bin_buf_Data.Bytes()))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)

	bin_buf_without_crc.Write(bin_buf_Data.Bytes())
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func __询查单张标签_EPC_C1_G2_ISO18000() (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {
	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}
		/*
			if element.C == command["询查多标签_EPC_C1G2"].C {
		*/
		if response.Status == 0x01 {
			log.Printf("成功读取  %x 个标签", b[:1])
			//log.Printf("成功读取  %d byte,读取内容:%x", nr, b[:nr])

			buff := bytes.NewBuffer(b)

			//re := bytes.NewReader(b);
			buf_count := make([]byte, 1)
			_, _ = buff.Read(buf_count)

			count := int(buf_count[0])

			//返回未读取部分的长度
			//	fmt.Println("re len : ", re.Len());
			//返回底层数据总长度
			//	fmt.Println("re size : ", re.Size());

			epc_ids := make([][]byte, count)
			for i := 0; i < count; i++ {
				buf_epc_len := make([]byte, 1)
				buff.Read(buf_epc_len)
				len_epc := int(buf_epc_len[0])

				buf_epc_code := make([]byte, len_epc)
				buff.Read(buf_epc_code)

				log.Printf("读取一个  %x 个标签", i)

				log.Printf("读取一个 长度：%d 的标签 %x ", len_epc, buf_epc_code)
				epc_ids[i] = buf_epc_code

				//	fmt.Printf("----------- epc_ids_ %x\n", epc_ids[i])

				//	data_byte, _,element := 读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i])
				//	Send_(data_byte,element)
				//	fmt.Printf("----------- epc_ids_ %x", data_byte)

				//	data_byte, _,element := __写EPC号_EPC_C1_G2_ISO18000([4]byte{0x01,0x01,0x01,0x01},[]byte{0x01,0x01,0x01,0x01})

				//	data_byte, fun,element := _______读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i])
				data_byte, fun, element := __读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i], Policy_TID存储区, 0x00, 0x08)
				Send_(data_byte, element, fun)

				fmt.Printf("----------- epc_ids_ %x", data_byte)

			}

			return epc_ids, err

		}

		return nil, err
	}

	var dat = []uint8{} //[]uint8{0x05}

	element = command["询查单张标签_EPC_C1_G2_ISO18000-6C"]
	var cmd = command["询查单张标签_EPC_C1_G2_ISO18000-6C"].C

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

func __设定存储区读写保护状态_EPC_C1_G2_ISO18000(EPC []byte, Select Select_mode, SetProtect Select_SetProtect_mode, Pwd [4]byte) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {
	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}
		/*
			if element.C == command["询查多标签_EPC_C1G2"].C {
		*/
		if response.Status == 0x00 {
			log.Printf("设置成功 成功读取  %x 个标签", b[:1])

			return epc_ids, err

		}

		return nil, err
	}

	var data Data设定存储区读写保护状态_data
	data.Pwd = Pwd
	data.Select = Select_mode.String(Select)
	data.SetProtect = Select_SetProtect_mode.String(SetProtect)

	data.MaskLen = 0x00
	data.MaskAdr = 0x00

	var bin_buf_Data设定存储区读写保护状态_data bytes.Buffer
	binary.Write(&bin_buf_Data设定存储区读写保护状态_data, binary.BigEndian, data)

	element = command["设定存储区读写保护状态"]
	var cmd = element.C
	var length = uint8(4 + 1 + len(EPC) + 8)

	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)

	var NEPC_buf bytes.Buffer
	binary.Write(&NEPC_buf, binary.BigEndian, uint8(len(EPC)/2))
	bin_buf_without_crc.Write(NEPC_buf.Bytes())
	bin_buf_without_crc.Write(EPC)
	bin_buf_without_crc.Write(bin_buf_Data设定存储区读写保护状态_data.Bytes())

	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func 读数据_18000_6B() (bb []byte, work func(client uint16, msg uint16), element commandElement) {
	var f = func(client uint16, msg uint16) {
	}
	var dat = []uint8{} //[]uint8{0x05}

	element = command["读数据_18000-6B"]
	var cmd = command["读数据_18000-6B"].C

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

func 工作模式设置命令(Read_mode Read_mode_工作模式, Mode_state Mode_state_工作模式) (bb []byte, work func(client uint16, msg uint16), element commandElement) {
	var f = func(client uint16, msg uint16) {
	}

	var dat Data工作模式设置命令_data

	element = command["读取工作模式参数"]
	dat.Read_mode = Read_mode_工作模式.String(Read_mode)
	dat.Mode_state = Mode_state_工作模式.String(Mode_state)
	dat.Mem_Inven = 0x01
	dat.First_Adr = 0x01
	dat.Word_Num = 0x01
	dat.Tag_Time = 0x01
	var bin_buf_dat bytes.Buffer
	binary.Write(&bin_buf_dat, binary.BigEndian, dat)

	var length = uint8(4 + unsafe.Sizeof(dat))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = element.C
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(bin_buf_dat.Bytes())
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func EAS检测精度设置(Accuracy uint8) (bb []byte, work func(client uint16, msg uint16), element commandElement) {

	var f = func(client uint16, msg uint16) {

	}

	element = command["EAS检测精度设置"]
	var cmd = command["EAS检测精度设置"].C
	var dat = []uint8{0x10} //[]uint8{0x05}
	dat[0] = Accuracy

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

func __测试标签是否被设置读保护() (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {

		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		var by = b[:1]
		if response.Status == 0 {

			//	log.Printf("buf____4: %x", b)
			/*			ReadPro	说明
						0x00	电子标签没有被设置为读保护。
						0x01	电子标签被设置读保护。
			*/
			log.Printf("其他事情，")

			log.Printf("读取内容 %#x", b)
			log.Printf("读取内容 ReadPro%#x", by)
			if by[0] == 0x00 {
				log.Printf("电子标签没有被设置为读保护 %#x", by)
			} else {
				log.Printf("电子标签被设置读保护 %#x", by)

			}

		} else {

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err

	}

	element = command["测试标签是否被设置读保护"]
	var cmd = element.C
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

func __读保护设置_根据EPC号设定(Pwd [4]byte, EPC []byte) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	/*	Data[]
		ENum	EPC	Pwd	MaskAdr	MaskLen
		0xXX	变长	4Byte	0xXX	0xXX*/

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD || response.Status == 0x0C {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			log.Printf("__读保护设置_根据EPC号设定 成功，")

		} else {
			log.Printf("__读保护设置_失败，")

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err
	}

	element = command["读保护设置(根据EPC号设定)"]
	var cmd = element.C

	var bin_buf_Data bytes.Buffer
	var ENum uint8 = uint8(len(EPC) / 2)
	eNum := make([]byte, 1)
	eNum[0] = ENum
	bin_buf_Data.Write(eNum)
	bin_buf_Data.Write(EPC)
	bin_buf_Data.Write(Pwd[:])

	/*	bin_buf_Data.Write([]byte{0x00})
		bin_buf_Data.Write([]byte{0x01})*/

	var length = uint8(4 + len(bin_buf_Data.Bytes()))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(bin_buf_Data.Bytes())
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func __解锁读保护(Pwd [4]byte, EPC []byte) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	/*	Data[]
		ENum	EPC	Pwd	MaskAdr	MaskLen
		0xXX	变长	4Byte	0xXX	0xXX*/

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			log.Printf("__读保护设置_根据EPC号设定 成功，")

		} else {
			log.Printf("__读保护设置_失败，")

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err
	}

	element = command["解锁读保护"]
	var cmd = element.C

	var bin_buf_Data bytes.Buffer
	/*	var ENum = uint8(len(EPC)/2)
		eNum := make([]byte, 1)
		eNum[0] = ENum
		bin_buf_Data.Write(eNum)
		bin_buf_Data.Write(EPC)*/
	bin_buf_Data.Write(Pwd[:])

	/*	bin_buf_Data.Write([]byte{0x00})
		bin_buf_Data.Write([]byte{0x01})
	*/

	var length = uint8(4 + len(bin_buf_Data.Bytes()))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(bin_buf_Data.Bytes())
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}

func __写数据(Pwd [4]byte, EPC []byte, Mem PolicyType, WordPtr uint8, Wdt []byte) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	/*	Data[]
		ENum	EPC	Pwd	MaskAdr	MaskLen
		0xXX	变长	4Byte	0xXX	0xXX*/

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD || response.Status == 0x0C {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			log.Printf("____写数据 成功，")

		} else {
			log.Printf("____写数据_失败，")

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err
	}

	element = command["写数据"]
	var cmd = element.C

	var bin_buf_Data bytes.Buffer

	var WNum uint8 = uint8(len(Wdt) / 2)
	wNum := make([]byte, 1)
	wNum[0] = WNum
	bin_buf_Data.Write(wNum)

	var ENum uint8 = uint8(len(EPC) / 2)
	eNum := make([]byte, 1)
	eNum[0] = ENum
	bin_buf_Data.Write(eNum)

	bin_buf_Data.Write(EPC)

	mem := make([]byte, 1)
	mem[0] = PolicyType.String(Mem)
	bin_buf_Data.Write(mem)

	wordPtr := make([]byte, 1)
	wordPtr[0] = WordPtr
	bin_buf_Data.Write(wordPtr)

	bin_buf_Data.Write(Wdt)
	bin_buf_Data.Write(Pwd[:])

	bin_buf_Data.Write([]byte{0x00})
	bin_buf_Data.Write([]byte{0x01})

	var length = uint8(4 + len(bin_buf_Data.Bytes()))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(bin_buf_Data.Bytes())
	checksum := crc16.Checksum(bin_buf_without_crc.Bytes(), crc16.MakeTable(crc16.CRC16_MCRF4XX))
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, checksum)
	bin_buf_without_crc.Write(b)
	return bin_buf_without_crc.Bytes(), f, element

}
func __块擦除(Pwd [4]byte, EPC []byte, Mem PolicyType, WordPtr uint8, Num uint8) (bb []byte, work func(response Reponse, b []byte) (epc_ids [][]byte, err error), element commandElement) {

	/*	Data[]
		ENum	EPC	Pwd	MaskAdr	MaskLen
		0xXX	变长	4Byte	0xXX	0xXX*/

	var f = func(response Reponse, b []byte) (epc_ids [][]byte, err error) {
		if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status == 0xFD || response.Status == 0x0C {
			log.Printf("命令执行异常 %s", StatusMap[response.Status])
		}

		if response.Status == 0 {

			log.Printf("____写数据 成功，")

		} else {
			log.Printf("____写数据_失败，")

			if response.Status == 0xFE || response.Status == 0xfb {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])

			}

		}

		return nil, err
	}

	element = command["块擦除"]
	var cmd = element.C

	var bin_buf_Data bytes.Buffer

	var ENum uint8 = uint8(len(EPC) / 2)
	eNum := make([]byte, 1)
	eNum[0] = ENum
	bin_buf_Data.Write(eNum)

	bin_buf_Data.Write(EPC)

	mem := make([]byte, 1)
	mem[0] = PolicyType.String(Mem)
	bin_buf_Data.Write(mem)

	wordPtr := make([]byte, 1)
	wordPtr[0] = WordPtr
	bin_buf_Data.Write(wordPtr)

	num := make([]byte, 1)
	num[0] = Num
	bin_buf_Data.Write(num)

	bin_buf_Data.Write(Pwd[:])

	bin_buf_Data.Write([]byte{0x00})
	bin_buf_Data.Write([]byte{0x01})

	var length = uint8(4 + len(bin_buf_Data.Bytes()))
	header_without_crc := &Header_without_crc{Address: 0x00}
	header_without_crc.Command = cmd
	header_without_crc.Length = length
	var bin_buf_without_crc bytes.Buffer
	binary.Write(&bin_buf_without_crc, binary.BigEndian, header_without_crc)
	bin_buf_without_crc.Write(bin_buf_Data.Bytes())
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

	bin_buf_without_crc, _, element := __询查多标签_EPC_C1G2(0xFF, 0xFF)

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

		if element.C == command["询查多标签_EPC_C1G2"].C {

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

func InitRfid(epc_id [][]byte) (name Peripheral, err error) {
	//设置串口编号
	Name := "rfid COM31"
	c := &serial.Config{Name: "COM31", Baud: 57600, ReadTimeout: time.Second * 2}
	//打开串口
	s, err = serial.OpenPort(c)

	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

	return Peripheral{Name: Name}, err
}

func Send_(bin_buf_without_crc []byte, element commandElement, work func(response Reponse, b []byte) (epc_ids [][]byte, err error)) (err error) {

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

		if nr == 0 {
			break
		}
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

		work(response, b)

		/*		if element.C == command["读数据_EPC_C1_G2_ISO18000-6C"].C {

					if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status ==0xFD {
						log.Printf("命令执行异常 %s", StatusMap[response.Status])
					}



					if response.Status == 0x01 {

						buff :=bytes.NewBuffer(b)

						fmt.Printf("成功读取串口返回数据  %x",buff.Bytes())











						return  err





					}

					return  err
				}
		*/
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

	return err

}

func Send() (epc_ids_ [][]byte, err error) {

	//	bin_buf_without_crc,_,element := 读取工作模式参数()
	//bin_buf_without_crc,_,element := 工作模式设置命令(应答模式 ,输出方式选择位)

	//	bin_buf_without_crc,_,element := EAS检测精度设置(0x08)

	//	bin_buf_without_crc, fun, element := __询查单张标签_EPC_C1_G2_ISO18000()
	//	bin_buf_without_crc, fun, element := A声光控制命令(0x10,0x10,0x01)
	//bin_buf_without_crc, fun, element := __读取读写器信息()
	//bin_buf_without_crc, fun, element := ______写EPC号_EPC_C1_G2_ISO18000([4]byte{0x00,0x00,0x00,0x00},[]byte{0x01,0x01,0x01,0x01})

	//bin_buf_without_crc, _, element := __询查多标签_EPC_C1G2(0xFF, 0xFF)

	//bin_buf_without_crc, fun, element := __询查多标签_EPC_C1G2(0xFF, 0xFF)
	//	bin_buf_without_crc, fun, element := __测试标签是否被设置读保护()
	//bin_buf_without_crc, fun, element := __读保护设置_根据EPC号设定([4]byte{0x00,0x00,0x00,0x00},[]byte{0x01,0x01,0x01,0x01})

	bin_buf_without_crc, fun, element := __读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2([]byte{0x01, 0x01, 0x01, 0x01}, Policy_EPC存储区, 0x00, 0x08)
	//([]byte{0x01,0x01,0x01,0x01}, Policy_TID存储区,0x00, 0x0c)  =13*16 = 208bit
	//([]byte{0x01,0x01,0x01,0x01}, Policy_EPC存储区,0x00, 0x06) （2+6）*16 = 32+96bit =128bit
	//([]byte{0x01,0x01,0x01,0x01}, Policy_保留区,0x00, 0x0f) =4*16 =64
	//([]byte{0x01,0x01,0x01,0x01}, Policy_用户存储区,0x00, 0x1f) =32 * 16=512
	//256位EPC码，512位用户数据区，96位TID码
	//bin_buf_without_crc, _, element := 询查多标签_EPC_C1G2(0xFF, 0xFF)

	//var Wdt = []byte{0x01,0x02,0x03,0x04,0x05,0x06,0x07,0x08}
	//bin_buf_without_crc, fun,element := __写数据([4]byte{0x00,0x00,0x00,0x00},[]byte{0x01,0x01,0x01,0x01}, Policy_用户存储区,0x00, Wdt)
	//                          func __写数据(Pwd [4]byte,EPC []byte, Mem PolicyType,WordPtr uint8,WNum uint8, Wdt uint8) (bb []byte, work func(response Reponse, b []byte)(epc_ids [][]byte,err error), element commandElement) {

	//bin_buf_without_crc, fun,element := __块擦除([4]byte{0x00,0x00,0x00,0x00},[]byte{0x01,0x01,0x01,0x01}, Policy_用户存储区,0x00, 0x04)

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

		if nr == 0 {
			break
		}
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

		return fun(response, b)

		/*	if response.Status == 0xFE || response.Status == 0xfb || response.Status == 0xff || response.Status ==0xFD {
				log.Printf("命令执行异常 %s", StatusMap[response.Status])
			}


			if element.C == command["询查多标签_EPC_C1G2"].C {

				if response.Status == 0x01 {
					log.Printf("成功读取  %x 个标签", b[:1])
					//log.Printf("成功读取  %d byte,读取内容:%x", nr, b[:nr])

					buff :=bytes.NewBuffer(b)


					//re := bytes.NewReader(b);
					buf_count := make([]byte, 1);
					_, _=buff.Read(buf_count)



					count := int(buf_count[0])



					//返回未读取部分的长度
					//	fmt.Println("re len : ", re.Len());
					//返回底层数据总长度
					//	fmt.Println("re size : ", re.Size());

					epc_ids := make([][]byte,count)
					for i = 0; i < count; i++ {
						buf_epc_len := make([]byte, 1);
						buff.Read(buf_epc_len)
						len_epc := int(buf_epc_len[0])

						buf_epc_code := make([]byte, len_epc);
						buff.Read(buf_epc_code)


						log.Printf("读取一个  %x 个标签", i)

						log.Printf("读取一个 长度：%d 的标签 %x ", len_epc,buf_epc_code)
						epc_ids[i] = buf_epc_code

					//	fmt.Printf("----------- epc_ids_ %x\n", epc_ids[i])


					//	data_byte, _,element := 读数据_EPC_C1_G2_ISO18000_6C_EPC_C1G2(epc_ids[i])
					//	Send_(data_byte,element)
					//	fmt.Printf("----------- epc_ids_ %x", data_byte)



						data_byte, _,element := 写EPC号_EPC_C1_G2_ISO18000(1,0,[]byte{0x01,0x01,0x01,0x01})
							Send_(data_byte,element)
						fmt.Printf("----------- epc_ids_ %x", data_byte)
					}











					return epc_ids, err





				}

				return nil, err
			}*/

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

	return nil, err

}

func readNextBytes(file *os.File, number int) []byte {
	bytes := make([]byte, number)

	_, err := file.Read(bytes)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

var StatusMap = map[uint8]string{
	0xFE: "不合法的命令	当上位机输入的命令是不可识别的命令，如不存在的命令、或是CRC错误的命令",
	0xfb: "无电子标签可操作	当读写器对电子标签进行操作时，有效范围内没有可操作的电子标签时返回给上位机的状态值",
	0xFF: "参数错误	上位机发送的命令中的参数不符合要求时，返回此状态",
	0xFD: "命令长度错误	当上位机输入的命令的实际长度和它应当具有的长度不同时，返回该状态",
	0xFA: "无此项	LSB+MSB	有电子标签，但通信不畅，操作失败	当检测到有效范围内存在可操作的电子标签，但读写器与电子标签之间的通讯质量不好，而无法完成整个通讯过程时返回给上位机的信息",
	0xFC: "电子标签返回错误代码	电子标签返回错误代码时，错误代码由Err_code返回给上位机",
	0x0C: "对该命令访问密码不能为全0	对NXP UCODE EPC G2X标签设置读保护及设置EAS报警时，访问密码不能为全0，若为全0，将返回此状态值",
	0x05: "访问密码错误	当读写器执行需要密码才能执行的操作，而命令中给出的密码是错误的密码时返回给上位机的状态值",
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
