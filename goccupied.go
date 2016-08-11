package main

import (
	"github.com/evodicka/goccupied/handler"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	go func() {
		sigchan := make(chan os.Signal, 10)
		signal.Notify(sigchan, os.Interrupt, os.Kill)
		<-sigchan
		handler.Close()

		os.Exit(0)
	}()

	handler.Init()
	defer handler.Close()
	handler.Starting()

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/occupied", handler.GetOccupiedState).Methods("GET")
	router.HandleFunc("/api/v1/toggle", handler.ToggleLed).Methods("POST")

	handler.Started()
	http.ListenAndServe(":8000", router)
}
