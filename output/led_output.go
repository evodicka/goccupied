package output

import (
	gpio "github.com/stianeikeland/go-rpio"
)

func ToggleLed() {
	gpio.Open()
	defer gpio.Close()

	pin := gpio.Pin(7)
	pin.Output()
	pin.Toggle()
}
