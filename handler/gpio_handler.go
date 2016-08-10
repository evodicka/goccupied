package handler

import (
	"encoding/json"
	"github.com/evodicka/goccupied/output"
	"github.com/evodicka/goccupied/sensor"
	"net/http"
	"strconv"
)

// Handler function to read the occupied state from a sensor. Writes a json object containing the state
// into the response
func GetOccupiedState(rw http.ResponseWriter, r *http.Request) {

	light := sensor.LightOn()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(light)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseBody)

}

// Handler function to switch a LED on and off
func ToggleLed(rw http.ResponseWriter, r *http.Request) {
	output.ToggleLed()

	rw.WriteHeader(http.StatusOK)
}
