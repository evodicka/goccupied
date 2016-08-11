package output

import (
	"github.com/stianeikeland/go-rpio"
	"time"
)

var (
	busyPin   = rpio.Pin(10)
	activePin = rpio.Pin(9)
)

// Initializes the LED output pins
func InitLEDs() {
	busyPin.Output()
	activePin.Output()
}

// Blinks n times on the busy pin
func Blink(times int) {
	// Toggle pin 20 times
	for x := 0; x < times; x++ {
		busyPin.Toggle()
		time.Sleep(time.Second / 5)
	}
}

// Switches the state of GPIO10
func ToggleBusy() {
	busyPin.Toggle()
}

func SwitchBusyOn() {
	busyPin.High()
}

func SwitchBusyOff() {
	busyPin.Low()
}

func SwitchActiveOn() {
	activePin.High()
}

func SwitchActiveOff() {
	activePin.Low()
}
