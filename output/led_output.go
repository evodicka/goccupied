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

// Switches GPIO10 on
func SwitchBusyOn() {
	busyPin.High()
}

// Switches GPIO10 off
func SwitchBusyOff() {
	busyPin.Low()
}

// Switches GPIO09 on
func SwitchActiveOn() {
	activePin.High()
}

// Switches GPIO09 off
func SwitchActiveOff() {
	activePin.Low()
}
