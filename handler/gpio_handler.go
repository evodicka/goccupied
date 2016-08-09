package handler

import (
	"encoding/json"
	"github.com/evodicka/goccupied/output"
	"github.com/evodicka/goccupied/sensor"
	"net/http"
	"strconv"
)

func GetOccupiedState(rw http.ResponseWriter, r *http.Request) {

	light := sensor.LightOn()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(light)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseBody)

}

func ToggleLed(rw http.ResponseWriter, r *http.Request) {
	output.ToggleLed()

	rw.WriteHeader(http.StatusOK)
}
