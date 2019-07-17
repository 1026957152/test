package hikvision

import (
	"test/src/drivers"
	"test/src/platforms/hik2"
	"time"
)

func main() {
	firmataAdaptor := hik2.NewAdaptor("/dev/ttyACM0")
	led := hik.NewLedDriver(firmataAdaptor, "13")

	work := func() {
		Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := NewRobot("bot",
		[]Connection{firmataAdaptor},
		[]Device{led},
		work,
	)

	robot.Start()
}
