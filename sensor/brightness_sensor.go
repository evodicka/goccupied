package sensor

import (
	"github.com/stianeikeland/go-rpio"
)

var (
	pin = rpio.Pin(14)
)

// Initializes the input pin for the sensor
func InitSensor() {
	pin.Input()
	pin.PullDown()
}

// Determines if the light is currently on. Reads from GPIO14. If the pin state is low, the sensor indicates that
// the light is on.
func IsLightOn() bool {
	state := pin.Read()
	return state == rpio.Low
}
