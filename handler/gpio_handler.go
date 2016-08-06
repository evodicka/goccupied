package handler

import (
	"encoding/json"
	"net/http"
	"github.com/evodicka/goccupied/sensor"
	"strconv"
	"github.com/evodicka/goccupied/output"
)

func GetOccupiedState(rw http.ResponseWriter, r *http.Request) {

	light := sensor.LightOn()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(light)

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(responseBody)

}

func ToggleLed(rw http.ResponseWriter, r *http.Request) {
	output.ToggleLed()

	rw.WriteHeader(http.StatusOK)
}
