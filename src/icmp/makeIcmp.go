---
contrib/external-tests/go/main.go | 72 ++++++++++++++++++++++++++++++++++-----
1 file changed, 64 insertions(+), 8 deletions(-)

diff --git a/contrib/external-tests/go/main.go b/contrib/external-tests/go/main.go
index 16632bb..68447fe 100644
--- a/contrib/external-tests/go/main.go
+++ b/contrib/external-tests/go/main.go
@@ -12,6 +12,8 @@ import (

"github.com/dchest/blake2s"
"github.com/titanous/noise"
+	"golang.org/x/net/icmp"
+	"golang.org/x/net/ipv4"
)

func main() {
@@ -36,6 +38,7 @@ func main() {
}
defer conn.Close()

+	// write handshake initiation packet
now := time.Now()
tai64n := make([]byte, 12)
binary.BigEndian.PutUint64(tai64n[:], uint64(now.Unix()))
@@ -53,6 +56,7 @@ func main() {
log.Fatalf("error writing initiation packet: %s", err)
}

+	// read handshake response packet
responsePacket := make([]byte, 89)
n, err := conn.Read(responsePacket)
if err != nil {
@@ -69,7 +73,7 @@ func main() {
if ourIndex != 28 {
log.Fatalf("response packet index wrong: want %d, got %d", 28, ourIndex)
}
-	payload, sendCipher, _, err := hs.ReadMessage(nil, responsePacket[9:57])
+	payload, sendCipher, receiveCipher, err := hs.ReadMessage(nil, responsePacket[9:57])
if err != nil {
log.Fatalf("error reading handshake message: %s", err)
}
@@ -77,12 +81,64 @@ func main() {
log.Fatalf("unexpected payload: %x", payload)
}

-	keepalivePacket := make([]byte, 13)
-	keepalivePacket[0] = 4 // Type: Data
-	binary.LittleEndian.PutUint32(keepalivePacket[1:], theirIndex)
-	binary.LittleEndian.PutUint64(keepalivePacket[5:], 0) // Nonce
-	keepalivePacket = sendCipher.Encrypt(keepalivePacket, nil, nil)
-	if _, err := conn.Write(keepalivePacket); err != nil {
-		log.Fatalf("error writing keepalive packet: %s", err)
+	// write ICMP Echo packet
+	pingMessage, _ := (&icmp.Message{
+		Type: ipv4.ICMPTypeEcho,
+		Body: &icmp.Echo{
+			ID:   1,
+			Seq:  1,
+			Data: []byte("WireGuard"),
+		},
+	}).Marshal(nil)
+	pingHeader, err := (&ipv4.Header{
+		Version:  ipv4.Version,
+		Len:      ipv4.HeaderLen,
+		TotalLen: ipv4.HeaderLen + len(pingMessage),
+		Protocol: 1, // ICMP
+		TTL:      2,
+		Checksum: 0xa15b, // the packet is always the same, hard-code checksum
+		Src:      net.IPv4(10, 189, 129, 2),
+		Dst:      net.IPv4(10, 189, 129, 1),
+	}).Marshal()
+	binary.BigEndian.PutUint16(pingHeader[2:], uint16(ipv4.HeaderLen+len(pingMessage))) // fix the length endianness on BSDs
+	if err != nil {
+		panic(err)
+	}
+	pingPacket := make([]byte, 13)
+	pingPacket[0] = 4 // Type: Data
+	binary.LittleEndian.PutUint32(pingPacket[1:], theirIndex)
+	binary.LittleEndian.PutUint64(pingPacket[5:], 0) // Nonce
+	pingPacket = sendCipher.Encrypt(pingPacket, nil, append(pingHeader, pingMessage...))
+	if _, err := conn.Write(pingPacket); err != nil {
+		log.Fatalf("error writing ping message: %s", err)
+	}
+
+	// read ICMP Echo Reply packet
+	replyPacket := make([]byte, 128)
+	n, err = conn.Read(replyPacket)
+	if err != nil {
+		log.Fatalf("error reading ping reply message: %s", err)
+	}
+	replyPacket = replyPacket[:n]
+	if replyPacket[0] != 4 { // Type: Data
+		log.Fatalf("unexpected reply packet type: %d", replyPacket[0])
+	}
+	replyPacket, err = receiveCipher.Decrypt(nil, nil, replyPacket[13:])
+	if err != nil {
+		log.Fatalf("error decrypting reply packet: %s", err)
+	}
+	replyHeaderLen := int(replyPacket[0]&0x0f) << 2
+	replyLen := binary.BigEndian.Uint16(replyPacket[2:])
+	replyMessage, err := icmp.ParseMessage(1, replyPacket[replyHeaderLen:replyLen])
+	if err != nil {
+		log.Fatalf("error parsing echo: %s", err)
+	}
+	echo, ok := replyMessage.Body.(*icmp.Echo)
+	if !ok {
+		log.Fatalf("unexpected reply body type %T", replyMessage.Body)
+	}
+
+	if echo.ID != 1 || echo.Seq != 1 || string(echo.Data) != "WireGuard" {
+		log.Fatalf("incorrect echo response: %#v", echo)
}
}
--
2.9.0

