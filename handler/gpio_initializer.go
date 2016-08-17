package handler

import (
	"github.com/evodicka/goccupied-client/output"
	"github.com/evodicka/goccupied-server/sensor"
	"github.com/stianeikeland/go-rpio"
)

// Initializes GPIO
func Init() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}
	sensor.InitSensor()
	output.InitLEDs()
}

// Switches the state to Starting
func Starting() {
	output.Blink(20)
}

// Switches the State to started
func Started() {
	output.SwitchBusyOff()
	output.SwitchFreeOn()
}

// Switches the state to Work started
func WorkStart() {
	output.SwitchBusyOn()
}

// Switches the state to work finished
func WorkStop() {
	output.SwitchBusyOff()
}

// Deactivates all Leds and closes GPIO
func Close() {
	output.SwitchBusyOff()
	output.SwitchFreeOff()
	rpio.Close()
}
