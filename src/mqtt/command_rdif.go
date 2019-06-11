package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"test/src/serial"

	"encoding/hex"
)

func command_rfid(client MQTT.Client, session_key_id string, message DownlinkMessageCommand) {

	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appId)
	link_topic := r.Replace(UPLINK_MESSAGE_t_down_acks)

	fmt.Printf("MSG: %s\n", message)

	if message.Command == "rfid" {
		fmt.Printf("射频命令 \n")
		if message.Confirmed == true {

			// step 下载文件
			// setp 更新docker
			correlation_ids := message.Correlation_ids
			uplink_confirm_mqtt(client, link_topic, correlation_ids) //an acknowledgement of a confirmed downlink

			cf, err := serial.Send()
			if err == nil {
				uplink_Messages_t_up_topic := r.Replace(Uplink_Messages_t_up)

				//	var uplink_message map[string]interface{} = make(map[string]interface{})

				var uplinkMessage UplinkMessage

				uplinkMessage.Session_key_id = session_key_id
				uplinkMessage.Uplink_token = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"
				encodedStr := hex.EncodeToString(cf)
				uplinkMessage.Pay_load = encodedStr // string(cf)
				//uplink_message[""] = cf
				uplink_Messages_t_up_mqtt(client, uplink_Messages_t_up_topic, uplinkMessage) //an acknowledgement of a confirmed downlink

				fmt.Printf("----------- images %x", cf)

			}

		}

	}

}
