package handler

import (
	"encoding/json"
	"github.com/evodicka/goccupied/output"
	//"github.com/evodicka/goccupied/sensor"
	"net/http"
	"strconv"
	"github.com/evodicka/goccupied/cache"
	"time"
	"github.com/evodicka/goccupied/sensor"
)

// Handler function to read the occupied state from a sensor. Writes a json object containing the state
// into the response
func GetOccupiedState(rw http.ResponseWriter, r *http.Request) {
	state, since := cache.GetState()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(state)
	responseBody["since"] = strconv.FormatInt(int64(since / time.Second), 10)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseBody)
}

// Handler function to switch a LED on and off
func ToggleLed(rw http.ResponseWriter, r *http.Request) {
	output.ToggleBusy()
	rw.WriteHeader(http.StatusAccepted)
}

func JustBlink() {
	output.ToggleBusy()
}

func ReadOccupiedState() {
	light := sensor.IsLightOn()
	cache.StoreState(!light)
}
