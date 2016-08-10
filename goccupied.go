package main

import (
	"github.com/evodicka/goccupied/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/occupied", handler.GetOccupiedState).Methods("GET")
	router.HandleFunc("/api/v1/toggle", handler.ToggleLed).Methods("GET")

	http.ListenAndServe(":8000", router)
}
