package gpio

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)


const (
	GreenLedPin1 uint8 = 21
	GreenLedPin2 uint8 = 22
	GreenLedPin3 uint8 = 23

	BlueLedPin1 uint8 = 13
	BlueLedPin2 uint8 = 15
	BlueLedPin3 uint8 = 16

	RedLedPin1 uint8 = 10
	RedLedPin2 uint8 = 11
	RedLedPin3 uint8 = 12
)


type CallBackPubGreenLedStatus func(interface{})
var PubGreenLedStatus CallBackPubGreenLedStatus

func Start() {

	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println("gpio : Start()")
	rpio.Open()
	defer rpio.Close()

	go goDancingGreenLed()// subscribing routine
	go goDancingBlueLed() // subscribing routine
	go goStillRedLed()    // publishing routine

	blockForever()
}

func goDancingGreenLed() {

	fmt.Println("goroutine : goDancingGreenLed")
	var gledPin1, gledPin2, gledPin3 rpio.Pin
	var gledState1, gledState2, gledState3 rpio.State

	gledPin1 = rpio.Pin(GreenLedPin1)
	gledPin2 = rpio.Pin(GreenLedPin2)
	gledPin3 = rpio.Pin(GreenLedPin3)

	gledPin1.Output()
	gledPin2.Output()
	gledPin3.Output()

	for {

		gledMap := make(map[string]string)
		gledState1 = GetRandomOnOff()
		gledState2 = GetRandomOnOff()
		gledState3 = GetRandomOnOff()

		rpio.WritePin(gledPin1, gledState1)
		rpio.WritePin(gledPin2, gledState2)
		rpio.WritePin(gledPin3, gledState3)

		if gledPin1.Read() == rpio.High {
			gledMap["LED1"] = "ON"
		} else {
			gledMap["LED1"] = "OFF"
		}

		if gledPin2.Read() == rpio.High {
			gledMap["LED2"] = "ON"
		} else {
			gledMap["LED2"] = "OFF"
		}

		if gledPin3.Read() == rpio.High {
			gledMap["LED3"] = "ON"
		} else {
			gledMap["LED3"] = "OFF"
		}

		gledJSON, _ := json.Marshal(gledMap)

		fmt.Println("gpio : Calling PubGreenLedStatus ==> ", gledMap)
		PubGreenLedStatus(gledJSON)
		time.Sleep(time.Second * 7)
	}
}