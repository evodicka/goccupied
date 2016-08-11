package handler

import (
	"github.com/evodicka/goccupied/output"
	"github.com/evodicka/goccupied/sensor"
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
	output.SwitchActiveOn()
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
	output.SwitchActiveOff()
	rpio.Close()
}
