package main

import (
	"fmt"
	"github.com/kr/pretty"
	"sync"
	"time"
)

func main() {

	type Project struct {
		Id    int64  `json:"project_id"`
		Title string `json:"title"`
		Name  string `json:"name"`
		//	Data Data `json:"data"`
		//	Commits Commits `json:"commits"`
	}

	var x = Project{1, "aaa", "{5, 6}"}
	fmt.Printf("%# v", pretty.Formatter(x)) //It will print all struct details

	fmt.Printf("%# v", pretty.Formatter(x.Id)) //It will print component one by one.

	// Create Udev
	u := udev.Udev{}

	// Create new Device based on subsystem and sysname
	d := u.NewDeviceFromSubsystemSysname("mem", "zero")

	// Extract information
	fmt.Printf("Sysname:%v\n", d.Sysname())
	fmt.Printf("Syspath:%v\n", d.Syspath())
	fmt.Printf("Devpath:%v\n", d.Devpath())
	fmt.Printf("Devnode:%v\n", d.Devnode())
	fmt.Printf("Subsystem:%v\n", d.Subsystem())
	fmt.Printf("Devtype:%v\n", d.Devtype())
	fmt.Printf("Sysnum:%v\n", d.Sysnum())
	fmt.Printf("IsInitialized:%v\n", d.IsInitialized())
	fmt.Printf("Driver:%v\n", d.Driver())
	fmt.Printf("Properties:%v\n", d.Properties())

	// Use one of the iterators

	// Use one of the iterators
	it := d.PropertyIterator()
	it.Each(func(item interface{}) {
		kv := item.([]string)
		_ = fmt.Sprintf("Property:%v=%v\n", kv[0], kv[1])
	})

	m := u.NewMonitorFromNetlink("udev")

	// Add filters to monitor
	//m.FilterAddMatchSubsystemDevtype("block", "disk")
	//m.FilterAddMatchTag("systemd")

	// Create a context
	ctx, cancel := context.WithCancel(context.Background())

	// Start monitor goroutine and get receive channel
	ch, _ := m.DeviceChan(ctx)

	// WaitGroup for timers
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		fmt.Println("Event:==================================")

		fmt.Println("Started listening on channel")
		for d := range ch {

			fmt.Println("==================================")

			fmt.Println("Event:", d.Syspath(), d.Action())
			// Extract information
			fmt.Printf("Sysname:%v\n", d.Sysname())
			fmt.Printf("Syspath:%v\n", d.Syspath())
			fmt.Printf("Devpath:%v\n", d.Devpath())
			fmt.Printf("Devnode:%v\n", d.Devnode())
			fmt.Printf("Subsystem:%v\n", d.Subsystem())

			mt.Printf("Devtype:%v\n", d.Devtype())
			fmt.Printf("Sysnum:%v\n", d.Sysnum())
			fmt.Printf("IsInitialized:%v\n", d.IsInitialized())
			fmt.Printf("Driver:%v\n", d.Driver())

			// Use one of the iterators
			it := d.PropertyIterator()
			it.Each(func(item interface{}) {
				kv := item.([]string)
				_ = fmt.Sprintf("Property:%v=%v\n", kv[0], kv[1])

			})

			fmt.Printf("Properties:%+v\n", d.Properties())

		}
		fmt.Println("Channel closed")
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting timer to update filter")
		<-time.After(10 * time.Second)
		fmt.Println("Removing filter")
		m.FilterRemove()
		fmt.Println("Updating filter")
		m.FilterUpdate()
		wg.Done()
	}()

	go func() {
		fmt.Println("Starting timer to signal done")
		<-time.After(20 * time.Second)
		fmt.Println("Signalling done")
		cancel()
		wg.Done()

	}()
	wg.Wait()

}
