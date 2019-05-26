package mqtt

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

type LocalDownlinks struct {
	Session_key_id  string
	Command         string
	Confirmed       bool
	Correlation_ids []string
}
type Msg struct {
	Play_load string
}

var localHandler MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("MSG: %s\n", msg.Payload())
	//text := fmt.Sprintf("this is result msg #%d!", knt)
	knt++
	var localDownlinks LocalDownlinks

	json.Unmarshal(msg.Payload(), &localDownlinks)

	if localDownlinks.Command == "info" {
		msgs := new(Msg)

		messageJson, _ := json.Marshal(info)
		msgs.Play_load = string(messageJson)

		mesJson, _ := json.Marshal(msgs)

		token := client.Publish("nn/result", 0, false, string(mesJson))
		token.Wait()
	}
	if localDownlinks.Command == "chromeup" {
		msgs := new(Msg)

		messageJson, _ := json.Marshal(info)
		msgs.Play_load = string(messageJson)

		mesJson, _ := json.Marshal(msgs)

		token := client.Publish("nn/result", 0, false, string(mesJson))
		token.Wait()
	}
	if localDownlinks.Command == "chromedown" {
		msgs := new(Msg)

		messageJson, _ := json.Marshal(info)
		msgs.Play_load = string(messageJson)

		mesJson, _ := json.Marshal(msgs)

		token := client.Publish("nn/result", 0, false, string(mesJson))
		token.Wait()
	}

}
var info map[string]string

func Mqtt_local(info_ map[string]string) {
	info = info_
	knt = 0

	opts := MQTT.NewClientOptions().AddBroker("tcp://localhost:1883")
	opts.SetClientID("mac-go")
	opts.SetUsername("11")
	opts.SetPassword("11")
	opts.SetDefaultPublishHandler(f)

	Client = MQTT.NewClient(opts)
	if token := Client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := Client.Subscribe("plain/led/status/green", 0, localHandler); token.Wait() &&
		token.Error() != nil {

		fmt.Println(token.Error())
		os.Exit(1)
	}

	var message map[string]string = make(map[string]string)
	message[""] = "aa"

	tokenPub := Client.Publish("plain/led/status/green", 0, false, "aaaaaaaaaaaa")
	tokenPub.Wait()

	time.Sleep(3 * time.Second)

} //
