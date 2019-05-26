package mqtt

import (
	//"camera"
	"encoding/json"
	"execC"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strings"
	"update"
)

var SubscribeHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	topic := msg.Topic()
	split_topic := strings.Split(topic, "/")
	deviceId := split_topic[2]

	appID := split_topic[0]

	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appID)
	link_topic := r.Replace(UPLINK_MESSAGE_t_down_acks)

	fmt.Printf("我在 mqtt subscrible 中了啊啊啊: %s\n", msg.Payload())

	fmt.Printf("MSG: %s\n", msg.Payload())
	//text:= fmt.Sprintf("this is result msg #%d!", knt)
	knt++

	var downlinkMessage DownlinkMessage

	var message = downlinkMessage.Downlinks

	//var m map[string]interface{}
	json.Unmarshal(msg.Payload(), &message)

	fmt.Printf("MSG: %s\n", message)

	/*	if message.Command == "image" {
			fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
			//execC.Newcamera()
			camera.Newcamera()
		}

		if message.Command == "vidoe" {
			fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
			//execC.Newcamera()
			camera.Newcamera()
		}
	*/

	if message.Command == "camera" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		//execC.Newcamera()
	}

	if message.Command == "docker" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		//execC.NewClient()
	}

	if message.Command == "scream" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		//execC.Newcamera()
	}

	if message.Command == "chroma" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		//execC.Newcamera()
		execC.Chrome_on()
		execC.Chrome_off()
	}

	if message.Command == "update" {
		fmt.Printf("更新设备 \n")
		if message.Confirmed == true {
			correlation_ids := message.Correlation_ids
			uplink_confirm_mqtt(client, link_topic, correlation_ids) //an acknowledgement of a confirmed downlink
			//downlink := message.Downlinks
			//var a interface{}
			//var b string
			//a = downlink
			////b = a.(string)
			//fmt.Println(a, b)

			err := update.DownloadFile_("e:\\a.txt", "https://raw.githubusercontent.com/idreamsi/RadioHead/master/LICENSE")
			if err != nil {
				uplink_Messages_t_up_topic := r.Replace(Uplink_Messages_t_up)

				var uplink_message map[string]interface{} = make(map[string]interface{})
				uplink_message["uplink_message"] = uplink_message
				uplink_message["session_key_id"] = "AWiZpAyXrAfEkUNkBljRoA=="
				uplink_message["uplink_token"] = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"

				uplink_Messages_t_up_mqtt(client, uplink_Messages_t_up_topic, uplink_message) //an acknowledgement of a confirmed downlink
			}

		}

		//update.DownloadFile_("e:\\a.txt","https://raw.githubusercontent.com/idreamsi/RadioHead/master/LICENSE")

		//	uplink_confirm_mqtt(client, link_topic,"complete")

	}

	/*	token := client.Publish("nn/result", 0, false, text)
		token.Wait()*/
}
