package main

import (
	"fmt"
	"github.com/kidoman/embd"
	_ "github.com/kidoman/embd/host/rpi" // This loads the RPi driver
	"os"
	"os/signal"
	//	"time"
)

func main() {
	err2 := embd.InitGPIO()
	fmt.Println("Initialized, err:", err2)
	defer embd.CloseGPIO()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt, stopping services...\n")
			cleanup()
			cleanupDone <- true
		}
	}()

	switchOff()
	//	time.Sleep(365 * 24 * 3600 * time.Second)
	<-cleanupDone
}

func switchOff() {
	pin, err := embd.NewDigitalPin(24)
	fmt.Println("Pin initialized, err:", err)
	pin.SetDirection(embd.Out)
	err = pin.Write(embd.High)
	fmt.Println("value high, err:", err)
}

func cleanup() {
	embd.CloseGPIO()
}
