package mqtt

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log"
	"test"
	//	"log"
	"strings"
	//gpio "test/src"
	"time"

	//"github.com/user/raspi-go-iot/gpio"
	"os"
)

var Uplink_Messages = "<AppID>/devices/<DevID>/up"
var Downlink_Messages = "<AppID>/devices/<DevID>/down"

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

func main_mqtt_local() {

	knt = 0

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("mac-go")
	opts.SetUsername("11")
	opts.SetPassword("11")
	opts.SetDefaultPublishHandler(f)

	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("plain/led/status/green", 0, nil); token.Wait() &&
		token.Error() != nil {

		fmt.Println(token.Error())
		os.Exit(1)
	}

	var message map[string]string = make(map[string]string)
	message[""] = "aa"

	tokenPub := c.Publish("plain/led/status/green", 0, false, "aaaaaaaaaaaa")
	tokenPub.Wait()

	time.Sleep(3 * time.Second)

} //
func New_mqtt(appID string, status map[string]string, server string) {

	var deviceId string = status["deviceEui"]
	log.Printf("9999999999999 %s-------%s", appID, deviceId)

	defer func() {
		fmt.Println("Mqqt 找值，defer end...")
	}()
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
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

	//
	r := strings.NewReplacer("<DevID>", deviceId, "<AppID>", appID)

	rt := r.Replace(Downlink_Messages)
	log.Printf("%s", rt)
	if token := c.Subscribe(rt, 0, subscribeHandler); token.Wait() &&
		token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)

	}

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

	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				// do stuff
				rt = r.Replace(Uplink_Messages)
				log.Printf("%s", rt)

				log.Printf("发送信息 %s", rt)
				var message map[string]string = make(map[string]string)
				message["aaa"] = "aa"
				messageJson, _ := json.Marshal(status)

				tokenPub := c.Publish(r.Replace(rt), 0, false, string(messageJson))
				tokenPub.Wait()
				log.Printf("完成发送信息 %s", rt)

			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

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

var subscribeHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

	fmt.Printf("我在 mqtt subscrible 中了啊啊啊: %s\n", msg.Payload())

	fmt.Printf("MSG: %s\n", msg.Payload())
	//text:= fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	var m map[string]string
	json.Unmarshal(msg.Payload(), &m)

	fmt.Printf("MSG: %s\n", m)

	if m["command"] == "camera" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		test.Newcamera()
	}

	if m["command"] == "docker" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		test.NewClient()
	}
	if m["command"] == "scream" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")
		test.Newcamera()
	}

	if m["command"] == "chroma" {
		fmt.Printf("LIGHT ON 打开灯啊啊啊\n")

		test.Newcamera()

		test.Chrome_on()
		test.Chrome_off()

	}
	/*	token := client.Publish("nn/result", 0, false, text)
		token.Wait()*/
}
