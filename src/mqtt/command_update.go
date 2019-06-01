package mqtt

import (
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"strings"
	"test/src/update"
)

func command_update(client MQTT.Client, message DownlinkMessage) {

	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appId)
	link_topic := r.Replace(UPLINK_MESSAGE_t_down_acks)

	fmt.Printf("MSG: %s\n", message)

	if message.Command == "update" {
		fmt.Printf("更新设备 \n")
		if message.Confirmed == true {

			// step 下载文件
			// setp 更新docker
			correlation_ids := message.Correlation_ids
			uplink_confirm_mqtt(client, link_topic, correlation_ids) //an acknowledgement of a confirmed downlink
			//downlink := message.Downlinks
			//var a interface{}
			//var b string
			//a = downlink
			////b = a.(string)
			//fmt.Println(a, b)

			url := "https://raw.githubusercontent.com/1026957152/test/master/src/dockerImages.yml?token=AAT2YJO6SOVPYF254IWKVY245PZSK"
			url = "https://raw.githubusercontent.com/1026957152/test/master/src/dockerImages.yml"
			//var filePath string = "~"+string(os.PathSeparator)+"a.txt"
			var filePath string = "e:\\" + string(os.PathSeparator) + "a.txt"

			cf, err := update.DownloadFile_(filePath, url)
			if err == nil {
				uplink_Messages_t_up_topic := r.Replace(Uplink_Messages_t_up)

				var uplink_message map[string]interface{} = make(map[string]interface{})
				uplink_message["session_key_id"] = "AWiZpAyXrAfEkUNkBljRoA=="
				uplink_message["uplink_token"] = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"

				uplink_Messages_t_up_mqtt(client, uplink_Messages_t_up_topic, uplink_message) //an acknowledgement of a confirmed downlink

				fmt.Printf("----------- images %s", cf.Images)

				//	docker.PullImage(cf.Images)

			}

			//docker.PullImage()
			//docker.RunContainer()
		}

		//update.DownloadFile_("e:\\a.txt","https://raw.githubusercontent.com/idreamsi/RadioHead/master/LICENSE")

		//	uplink_confirm_mqtt(client, link_topic,"complete")

	}

}
