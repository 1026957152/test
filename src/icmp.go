package main

import (
	"fmt"
	"golang.org/x/net/icmp"
	"net"
)

/*func main() {
	protocol := "icmp"
	netaddr, _ := net.ResolveIPAddr("ip4", "127.0.0.1")
	fmt.Printf("%s\n", netaddr)
	conn, _ := net.ListenIP("ip4:"+protocol, netaddr)
	fmt.Printf("net.ListenIP\n")

	buf := make([]byte, 1024)
	numRead, _, _ := conn.ReadFrom(buf)
	fmt.Printf("% X\n", buf[:numRead])
}
*/

func main() {
	netaddr, _ := net.ResolveIPAddr("ip4", "172.17.0.3")
	conn, _ := net.ListenIP("ip4:icmp", netaddr)
	for {
		buf := make([]byte, 1024)
		n, addr, _ := conn.ReadFrom(buf)
		msg,_:=icmp.ParseMessage(1,buf[0:n])

	//	m,_ := msg.Body.(*icmp.Echo)
		m,_ := msg.Body.(*icmp.Echo)

		fmt.Println(n, addr, msg.Type,msg.Code,msg.Checksum,m.,m.ID,m.Seq)
	}
}