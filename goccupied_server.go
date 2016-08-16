package main

import (
	"github.com/evodicka/goccupied-server/handler"
	"github.com/evodicka/goccupied-server/scheduler"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	go shutdown()

	handler.Init()
	defer handler.Close()
	handler.Starting()

	scheduler.Start(handler.ReadOccupiedState)

	router := mux.NewRouter()
	router.HandleFunc("/api/v1/occupied", handler.GetOccupiedState).Methods("GET")
	router.HandleFunc("/api/v1/light", handler.GetRawLightValue).Methods("GET")
	router.HandleFunc("/api/v1/toggle", handler.ToggleLed).Methods("POST")

	handler.Started()
	http.ListenAndServe(":8000", router)
}

func shutdown() {
	sigchan := make(chan os.Signal, 10)
	signal.Notify(sigchan, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-sigchan
	handler.Close()
	scheduler.Stop()

	os.Exit(0)
}
