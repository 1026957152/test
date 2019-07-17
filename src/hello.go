package main

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"os"
	"time"
)

//Product 商品信息
type Product struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

var knt int

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

func main() {

	var thingName string = "Led"
	var region string = "us-west-2"
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

	tokenPub := c.Publish("plain/led/status/green", 0, false, "aaaaaaaaaaaa")
	tokenPub.Wait()

	time.Sleep(3 * time.Second)

	p := &Product{}
	p.Name = thingName
	p.Name = region

	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	work := func() {
		mqttAdaptor.On("hello", func(msg mqtt.Message) {
			fmt.Println("hello")
		})
		mqttAdaptor.On("hola", func(msg mqtt.Message) {
			fmt.Println("hola")
		})
		data := []byte("o")
		gobot.Every(1*time.Second, func() {
			mqttAdaptor.Publish("hello", data)
		})
		gobot.Every(5*time.Second, func() {
			mqttAdaptor.Publish("hola", data)
		})
	}
} //
