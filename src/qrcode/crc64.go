package qrcode

import (
	"fmt"
	"hash/crc64"
)

func main() {
	s := "打死udhanckhdkja"
	//先创建一个table
	table := crc64.MakeTable(crc64.ECMA)
	//传入字节切片和table，返回一个uint64
	fmt.Println(crc64.Checksum([]byte(s), table)) //4295263180068867775

}
