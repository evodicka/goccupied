package handler

import (
	"encoding/json"
	"github.com/evodicka/goccupied-server/cache"
	"github.com/evodicka/goccupied-server/sensor"
	"net/http"
	"strconv"
	"time"
)

// Handler function to read the occupied state from a cache to indicate how long the current state allready exists.
// Writes a json object containing the state into the response
func GetOccupiedState(rw http.ResponseWriter, r *http.Request) {
	state, since := cache.GetState()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(state)
	responseBody["sinceSecconds"] = strconv.FormatInt(int64(since/time.Second), 10)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseBody)
}

// Returns the is light value directly from the sensor
func GetRawLightValue(rw http.ResponseWriter, r *http.Request) {
	light := sensor.IsLightOn()

	responseBody := make(map[string]string)
	responseBody["occupied"] = strconv.FormatBool(light)

	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(responseBody)
}

// Reads the current state from the sensor and writes it into the cache
func ReadOccupiedState() {
	light := sensor.IsLightOn()
	cache.StoreState(light)
}
