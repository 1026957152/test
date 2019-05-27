package mqtt

import (
	//"camera"
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	//"strings"
	"test/src/execC"
)

var SubscribeHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	/*	topic := msg.Topic()
		split_topic := strings.Split(topic, "/")
		deviceId := split_topic[2]

		appID := split_topic[0]

		r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appID)
		link_topic := r.Replace(UPLINK_MESSAGE_t_down_acks)*/

	fmt.Printf("我在 mqtt subscrible 中了啊啊啊: %s\n", msg.Payload())

	fmt.Printf("MSG: %s\n", msg.Payload())
	//text:= fmt.Sprintf("this is result msg #%d!", knt)
	knt++

	var downlink Downlink

	//var m map[string]interface{}
	json.Unmarshal(msg.Payload(), &downlink)

	var message DownlinkMessage
	json.Unmarshal([]byte(downlink.Downlinks.Pay_load), &message)

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

	command_update(client, message)

	/*	token := client.Publish("nn/result", 0, false, text)
		token.Wait()*/
}
