package mqtt

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	//"gopkg.in/matryer/try.v1"
	"log"
	//	"log"
	"strings"
	//gpio "test/src"
	"time"

	//"github.com/user/raspi-go-iot/gpio"
	"os"
)

var Uplink_Messages_t_up_image = "<AppID>/devices/<DevID>/up/image"
var Uplink_Messages_up_image = "<AppID>/devices/<DevID>/up/image"

var Uplink_Messages_t_up = "<AppID>/devices/<DevID>/up"

var UPLINK_MESSAGE_t_down_acks = "<AppID>/devices/<DevID>/down/acks"

var DOWNLINK_Messages_t_down = "<AppID>/devices/<DevID>/down"

type DownlinkMessageCommand struct {
	Session_key_id string
	Pay_load       string

	Command         string
	Confirmed       bool
	Correlation_ids []string
}

type DownlinkMessage struct {
	Session_key_id string
	Pay_load       string

	Command         string
	Confirmed       bool
	Correlation_ids []string
}

type Downlink struct {
	Downlinks DownlinkMessage
}

type UplinkMessage struct {
	Session_key_id  string   `json:"session_key_id"`
	Pay_load        string   `json:"pay_load"`
	Command         string   `json:"command"`
	Confirmed       bool     `json:"confirmed"`
	Correlation_ids []string `json:"correlation_ids"`
	Uplink_token    string   `json:"uplink_token"`
}
type Uplink struct {
	Uplink_message UplinkMessage `json:"uplink_message"`
	Device_id      string        `json:"device_id"`
	Application_id string        `json:"application_id"`
	Dev_eui        string        `json:"dev_eui"`
	Join_eui       string        `json:"join_eui"`
	Dev_addr       string        `json:"dev_addr"`

	/*	message["device_id"] = "dev1"
		message["application_id"] = "app1"
		message["dev_eui"] = "4200000000000000"
		message["join_eui"] = "4200000000000000"
		message["dev_addr"] = "01DA1F15"
		message["uplink_message"] = uplink_message*/
}

/*var Event_Messages = "<AppID>/devices/<DevID>/event"


downlink/scheduled
downlink/sent
activations
create
update
delete
down/acks
up/errors
down/errors
activations/errors

var Event_Messages = "<AppID>/devices/<DevID>/event"
*/

var knt int

/*func Init() {

	fmt.Println("pubsub : Raspberry Pi Pub/Sub initializing...")
	optsPub := MQTT.NewClientOptions()
	//optsPub.AddBroker("tcp://iot.eclipse.org:1883")
	optsPub.AddBroker("tcp://m14.cloudmqtt.com:14205")
	optsPub.SetUsername("htmbxcyz")     // TODO
	optsPub.SetPassword("rH2_IZj43nDy") // TODO
	optsPub.SetClientID("rasp-pi-go")
	optsPub.SetCleanSession(false)
	optsPub.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Println("SetDefaultPublishHandler : ", msg.Topic(), string(msg.Payload()))
	})

	fmt.Println("pubsub : Raspberry Pi MQTT broker configured")

	clientPub := MQTT.NewClient(optsPub)
	if tokenPub := clientPub.Connect(); tokenPub.Wait() && tokenPub.Error() != nil {
		panic(tokenPub.Error())
	}

/*	gpio.PubGreenLedStatus = func(ledMapJSON interface{}) {
		tokenPub := clientPub.Publish("plain/led/status/green", 0, false, ledMapJSON)
		tokenPub.Wait()
	}

	gpio.PubBlueLedStatus = func(ledMapJSON interface{}) {
		tokenPub := clientPub.Publish("secure/led/status/blue", 0, false, ledMapJSON)
		tokenPub.Wait()
	}

	gpio.PubRedLedStatus = func(ledMapJSON interface{}) {
		tokenPub := clientPub.Publish("secure/led/status/red", 0, false, ledMapJSON)
		tokenPub.Wait()
	}*/
/*
	if tokenPub := clientPub.Subscribe("secure/led/action/red", 0, gpio.SubRedLedAction); tokenPub.Wait() && tokenPub.Error() != nil {
		fmt.Println(tokenPub.Error())
		os.Exit(1)
	}*/

//	fmt.Println("pubsub : Raspberry Pi Pub/Sub callbacks registered")

//}

var Client MQTT.Client

var deviceId, appId string

func New_mqtt(appID string, status map[string]string, server string) {
	log.Printf("MAIN 主程序继续  New_mqtt")
	//	log.Printf("MAIN 主程序继续 New_mqtt", cf.AppID, status, cf.Server)

	deviceId = status["deviceEui"]
	appId = appID
	log.Printf("9999999999999 %s-------%s", appID, deviceId)

	defer func() {
		fmt.Println("Mqqt 错误 找值，defer end...")
	}()
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("MQTT 捕获到的错误：%s\n", r)
		}
	}()

	var thingName string = "Led"
	var region string = "us-west-2"

	opts := MQTT.NewClientOptions().AddBroker(server)
	opts.SetClientID("mac-go")
	opts.SetUsername("11")
	opts.SetPassword(region)
	opts.SetUsername(thingName)

	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	Client = c
	//
	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appID)

	rt := r.Replace(DOWNLINK_Messages_t_down)
	log.Printf("%s", rt)
	if token := c.Subscribe(rt, 0, SubscribeHandler); token.Wait() &&
		token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)

	}

	//v3/app1/devices/dev1/down

	//v3/{application id}/devices/{device id}/down/ack

	rt = r.Replace("<AppID>/devices/<DevID>/events/activations/errors")
	log.Printf("%s", rt)
	//Downlink Messages
	if token := c.Subscribe(rt, 0, nil); token.Wait() &&
		token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//Uplink Messages
	/*	tokenPub := c.Publish(r.Replace("<AppID>/devices/<DevID>/up"), 0, false, "aaaaaaaaaaaa")
		tokenPub.Wait()*/
	//Uplink Messages

	/*	ticker := time.NewTicker(10 * time.Second)
		quit := make(chan struct{})
		go func() {
			for {
				select {
				case <-ticker.C:
					// do stuff
					rt = r.Replace(Uplink_Messages_t_up)
					log.Printf("%s", rt)

					log.Printf("发送信息 %s", rt)
					var message map[string]interface{} = make(map[string]interface{})

					message["is_retry"] = false // Is set to true if this message is a retry (you could also detect this from the counter)
					message["confirmed"] = true

					message["latitude"] = 52.2345 //// Latitude of the device
					message["longitude"] = 6.2345  // Longitude of the device
					message["altitude"] = 2

					message["device_id"] = "dev1"
					message["application_id"] = "app1"
					message["dev_eui"] = "4200000000000000"
					message["join_eui"] = "4200000000000000"
					message["dev_addr"] = "01DA1F15"


					var uplink_message map[string]interface{} = make(map[string]interface{})
					message["uplink_message"] = uplink_message
					uplink_message["session_key_id"] = "AWiZpAyXrAfEkUNkBljRoA=="
					uplink_message["uplink_token"] = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"

				/*	"uplink_message": {
					"session_key_id": "AWiZpAyXrAfEkUNkBljRoA==",
						"f_port": 15,
						"frm_payload": "VGVtcGVyYXR1cmUgPSAwLjA=",
						"rx_metadata": [{
						"gateway_ids": {
							"gateway_id": "eui-0242020000247803",
							"eui": "0242020000247803"
						},
						"time": "2019-01-29T13:02:34.981Z",
						"timestamp": 1283325000,
						"rssi": -35,
						"snr": 5,
						"uplink_token": "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"
					}],



					messageJson, _ := json.Marshal(message)

					tokenPub := c.Publish(r.Replace(rt), 0, false, string(messageJson))
					tokenPub.Wait()
					log.Printf("完成发送信息 %s", rt)

				case <-quit:
					ticker.Stop()
					return
				}
			}
		}()
	*/
	//
	/*    tokenPub := c.Publish("<AppID>/devices/<DevID>/events/activations", 0, false, "aaaaaaaaaaaa")
	      tokenPub.Wait()*/

	log.Printf(" 结束mqtt 初始化")

	time.Sleep(3 * time.Second)

} //

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	token := client.Publish("nn/result", 0, false, text)
	token.Wait()
}

func register(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	token := client.Publish("nn/result", 0, false, text)
	token.Wait()
}

func uplink_confirm_mqtt(client MQTT.Client, uplink_acks string, correlation_ids interface{}) {
	//text := fmt.Sprintf("this is result msg #%d!", knt)
	fmt.Printf("发送了确认信息 acks \n")

	var message map[string]interface{} = make(map[string]interface{})

	message["is_retry"] = false // Is set to true if this message is a retry (you could also detect this from the counter)
	message["confirmed"] = true

	message["latitude"] = 52.2345 //// Latitude of the device
	message["longitude"] = 6.2345 // Longitude of the device
	message["altitude"] = 2

	message["device_id"] = deviceId
	message["application_id"] = "app1"
	message["dev_eui"] = deviceId
	message["join_eui"] = "4200000000000000"
	message["dev_addr"] = "01DA1F15"

	var downlink_ack map[string]interface{} = make(map[string]interface{})
	message["downlink_ack"] = downlink_ack
	downlink_ack["session_key_id"] = "AWiZpAyXrAfEkUNkBljRoA=="
	downlink_ack["uplink_token"] = "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"
	//	correlation_ids := [1]string{"my-correlation-id"}
	downlink_ack["correlation_ids"] = correlation_ids
	downlink_ack["confirmed"] = true
	downlink_ack["priority"] = "NORMAL"

	messageJson, _ := json.Marshal(message)
	knt++
	token := client.Publish(uplink_acks, 0, false, string(messageJson))
	token.Wait()
}
func uplink_Messages_t_up_mqtt(client MQTT.Client, uplink_acks string, uplink_message UplinkMessage) {

	//var message map[string]interface{} = make(map[string]interface{})
	var uplink Uplink
	uplink.Device_id = deviceId
	uplink.Application_id = "app1"
	uplink.Dev_addr = "01DA1F15"
	uplink.Dev_eui = deviceId
	uplink.Join_eui = "4200000000000000"
	uplink.Uplink_message = uplink_message

	/*	message["device_id"] = "dev1"
		message["application_id"] = "app1"
		message["dev_eui"] = "4200000000000000"
		message["join_eui"] = "4200000000000000"
		message["dev_addr"] = "01DA1F15"
		message["uplink_message"] = uplink_message*/

	messageJson, _ := json.Marshal(uplink)
	fmt.Printf("uplink_Messages_t_up_mqtt 发送了上传信息 %s \n", string(messageJson))

	//text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++

	//var value string
	/*	err := try.Do(func(attempt int) (bool, error) {
			var err error
		//	value, err = SomeFunction()
			token := client.Publish(uplink_acks, 0, false, string(messageJson),)
			token.Wait()
			return attempt < 5, err // try 5 times
		})
		if err != nil {
			log.Fatalln("error:", err)
		}*/
	token := client.Publish(uplink_acks, 0, false, string(messageJson))
	token.Wait()
}

func Uplink_Messages_t_up_mqtt(uplink_acks string, uplink_message map[string]interface{}) {

	var message map[string]interface{} = make(map[string]interface{})

	message["device_id"] = "dev1"
	message["application_id"] = appId
	message["dev_eui"] = deviceId
	message["join_eui"] = "4200000000000000"
	message["dev_addr"] = "01DA1F15"
	message["uplink_message"] = uplink_message
	fmt.Printf("发送了上传信息 \n")
	messageJson, _ := json.Marshal(message)

	//text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++

	//var value string
	/*	err := try.Do(func(attempt int) (bool, error) {
			var err error
		//	value, err = SomeFunction()
			token := client.Publish(uplink_acks, 0, false, string(messageJson),)
			token.Wait()
			return attempt < 5, err // try 5 times
		})
		if err != nil {
			log.Fatalln("error:", err)
		}*/
	token := Client.Publish(uplink_acks, 0, false, string(messageJson))
	token.Wait()
}

func Uplink_Messages_t_up_Image_mqtt(uplink_message []byte) {
	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appId)

	fmt.Printf("开始发送了上传信息 %s \n", r)

	token := Client.Publish(r.Replace(Uplink_Messages_t_up_image), 0, false, uplink_message)

	fmt.Printf("结束发送了上传信息 %s \n", r.Replace(Uplink_Messages_t_up_image))

	token.Wait()

}

/*				"uplink_message": {
				"session_key_id": "AWiZpAyXrAfEkUNkBljRoA==",
					"f_port": 15,
					"frm_payload": "VGVtcGVyYXR1cmUgPSAwLjA=",
					"rx_metadata": [{
					"gateway_ids": {
						"gateway_id": "eui-0242020000247803",
						"eui": "0242020000247803"
					},
					"time": "2019-01-29T13:02:34.981Z",
					"timestamp": 1283325000,
					"rssi": -35,
					"snr": 5,
					"uplink_token": "CiIKIAoUZXVpLTAyNDIwMjAwMDAyNDc4MDMSCAJCAgAAJHgDEMj49+ME"
				}],
*/
