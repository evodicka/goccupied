package output

import (
	gpio "github.com/stianeikeland/go-rpio"
)

// Switches the state of GPIO7
func ToggleLed() {
	gpio.Open()
	defer gpio.Close()

	pin := gpio.Pin(7)
	pin.Output()
	pin.Toggle()
}
