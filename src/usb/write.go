package main

import (
	//	"flag"
	"fmt"
	"log"

	"github.com/google/gousb"
	//	"github.com/google/gousb/usbid"
)

//var (
//	debugg = flag.Int("debug", 0, "libusb debug level (0..3)")
//)

func Write() {
	//	flag.Parse()

	// Initialize a new Context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	fmt.Println("--- 5 bytes successfully sent to the endpoint---- ")

	// Open any device with a given VID/PID using a convenience function.
	dev, err := ctx.OpenDeviceWithVIDPID(0x0525, 0xa4ac)

	//
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}

	fmt.Println("--- 5 bytes successfully sent to the endpoint---- ")

	fmt.Println("我的 设备 5 bytes successfully sent to the endpoint---- ")

	defer dev.Close()

	// The configurations can be examined from the DeviceDesc, though they can only
	// be set once the device is opened.  All configuration references must be closed,
	// to free up the memory in libusb.
	//               cfg,err := dev.Config(0)
	//if err != nil {
	//   log.Fatalf("%s.Config(2): %v", dev, err)
	//}
	//defer cfg.Close()

	for _, cfg := range dev.Desc.Configs { // This loop just uses more of the built-in and usbid pretty printing to list
		// the USB devices.
		fmt.Printf("  %s:\n", cfg)
		for _, intf := range cfg.Interfaces {
			fmt.Printf("    --------------\n")

			// Open an OUT endpoint.
			//  ep, err := intf.OutEndpoint(0x01)
			//    if err != nil {
			//              log.Fatalf("%s.OutEndpoint(7): %v", intf, err)
			//        }

			for _, ifSetting := range intf.AltSettings {
				fmt.Printf("if setting     %s\n", ifSetting)
				//                                        fmt.Printf("      %s\n", usbid.Classify(ifSetting))
				for _, end := range ifSetting.Endpoints {
					fmt.Printf("      %s\n", end)
				}
			}
		}
		fmt.Printf("    --------------\n")

	}
	// Claim the default interface using a convenience function.
	// The default interface is always #0 alt #0 in the currently active
	// config.
	log.Print("Enabling autodetach")
	dev.SetAutoDetach(true)

	cfg, err := dev.Config(1)
	if err != nil {
		log.Fatalf("%s.Config(2): %v", dev, err)
	}
	defer cfg.Close()

	// In the config #2, claim interface #3 with alt setting #0.
	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Fatalf("%s.Interface(0, 0): %v", cfg, err)
	}
	defer intf.Close()

	fmt.Printf("    CONFIG---------- %s-\n")

	// In this interface open endpoint #6 for reading.
	epIn, err := intf.InEndpoint(0x81)
	if err != nil {
		log.Fatalf("%s.InEndpoint(6): %v", intf, err)
	}

	fmt.Printf("    CONFIGIn Endpoint---------- %s-\n", epIn)

	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("%s.DefaultInterface(): %v", dev, err)
	}

	fmt.Println("我的 设备 5 bytes successfully sent to the endpoint---- ")

	defer done()

	// Open an OUT endpoint.
	ep, err := intf.OutEndpoint(7)
	if err != nil {
		log.Fatalf("%s.OutEndpoint(7): %v", intf, err)
	}

	// Generate some data to write.
	data := make([]byte, 5)
	for i := range data {
		data[i] = byte(i)
	}

	// Write data to the USB device.
	numBytes, err := ep.Write(data)
	if numBytes != 5 {
		log.Fatalf("%s.Write([5]): only %d bytes written, returned error is %v", ep, numBytes, err)
	}
	fmt.Println("5 bytes successfully sent to the endpoint")


	var e gousb.HotplugEvent


	ctx := gousb.NewContext()


	ctx.NewHotplug()

	hp := ctx.NewHotplug().Match(hotplug.Enter())
	for _, product := range products {
		hp.MatchAll(hotplug.Vendor(product.VendorID), hotplug.Product(product.ProductID))
	}
	hp.MatchAny(hotplug.Class(usb.HID), hotplug.Class(usb.MSC), hotplug.Class(usb.UMS))
	hp.Callback(func(ev *hotplug.Event) { ... })
	hp.CallbackOnce(func(ev *hotplug.Event) { ... })




	var cb HotplugCallback

	cb = ctx.Hotplug().Enumerate().ProductID(42).Register(func(event gousb.HotplugEvent) {
		cb.Deregister()
	})


	ctx.Hotplug.Enumerate().Register(func() {
		...
	})


	bus, addr := desc.Bus, desc.Address
	ctx.RegisterHotplug(func(ev *gosusb.HotplugEvent)) {
		if ev.DeviceDesc.Bus != bus || ev.DeviceDesc.Addr != addr || ev.Type != gousb.HotplugLeft {
			return
		}
		// do some cleanup now that the device has left
		ev.Cancel()
	})

	if (event.Type == gousb.HotplugEventDeviceArrived || event.Type == gousb.HotplugEventDeviceEnumerated) {
		// handle device arrived
	} else {
		// handle device left
	}
}
