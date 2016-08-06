package main

import (
	"github.com/evodicka/goccupied/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/occupied", handler.GetOccupiedState).Methods("GET")
	router.HandleFunc("/toggle", handler.ToggleLed).Methods("GET")

	http.ListenAndServe(":8000", router)
}
