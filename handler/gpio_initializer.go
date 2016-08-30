package handler

import (
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
}

// Deactivates all Leds and closes GPIO
func Close() {
	rpio.Close()
}
