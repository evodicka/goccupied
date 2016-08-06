package sensor

import (
	gpio "github.com/stianeikeland/go-rpio"
)

func LightOn() bool {
    gpio.Open()
	defer gpio.Close()

	pin := gpio.Pin(4)
	pin.Input()

	state := pin.Read()
	return state == gpio.Low
}
