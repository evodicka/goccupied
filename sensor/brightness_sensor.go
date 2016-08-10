package sensor

import (
	gpio "github.com/stianeikeland/go-rpio"
)

// Determines if the light is currently on. Reads from GPIO4. If the pin state is low, the sensor indicates that
// the light is on.
func LightOn() bool {
	gpio.Open()
	defer gpio.Close()

	pin := gpio.Pin(4)
	pin.Input()

	state := pin.Read()
	return state == gpio.Low
}
