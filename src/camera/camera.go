package camera

import (
	"fmt"
	"github.com/gpmgo/gopm/modules/log"
	//"gocv.io/x/gocv"

	//"gocv.io/x/gocv"
	"io/ioutil"
	"mqtt"
	"os"
)

var flag bool = false

func Newcamera() {

	if flag {
		return
	}
	/*	webcam, _ := gocv.VideoCaptureDevice(0)
		flag = true
		window := gocv.NewWindow("Hello")
		img := gocv.NewMat()

		for {
			webcam.Read(&img)
			window.IMShow(img)
			window.WaitKey(1)
		}*/
}

/*func main() {

	test := &Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Optionalgroup: &Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}

	log.Printf("Unmarshalled to: %+v", newTest)

}

*/
/*
func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	window := gocv.NewWindow("Hello")
	img := gocv.NewMat()

	for {
		webcam.Read(&img)
		window.IMShow(img)
		window.WaitKey(1)
	}
}
*/

var (
	mqttBroker   = "tcp://HOST:1883"
	mqttClientId = "Send_image_mqtt"
	mqttTopic    = "test/topic"
	imageFile    = "E:\\go\\src\\test\\src\\test.png"
)

func ImageMain() {

	/*	opts := MQTT.NewClientOptions()
		opts.AddBroker(mqttBroker)
		opts.SetClientID(mqttClientId)

		c := MQTT.NewClient(opts)
		if token := c.Connect(); token.Wait() && token.Error() != nil {
			log.Warn("Error connecting to broker: ", mqttBroker)
			panic(token.Error())
		}
	*/
	fileHandle, err := os.Open(imageFile)
	if err != nil {
		log.Fatal("Error opening file.")
	}

	defer func() {
		if err := fileHandle.Close(); err != nil {
			panic(err)
		}
	}()

	contents, err := ioutil.ReadAll(fileHandle)
	if err != nil {
		log.Fatal("Error reading file.")
	}

	fmt.Println(len(contents))

	mqtt.Uplink_Messages_t_up_Image_mqtt(contents)

	/*	token := c.Publish(mqttTopic, 0, false, string(contents))
		token.Wait()
		log.Info(token.Error())*/

}
