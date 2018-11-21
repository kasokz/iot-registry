package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Functionality struct, defining device functionality
type Functionality struct {
	Description string `json:"desc"`
	Path        string `json:"path"`
}

// Functionalities is a Functionality array
type Functionalities []Functionality

// Device structure
type Device struct {
	ID    string          `json:"id"`
	Name  string          `json:"name"`
	Funcs Functionalities `json:"funcs"`
}

// Devices is a Device array
type Devices []Device

func allDevices(w http.ResponseWriter, r *http.Request) {
	devices := Devices{
		Device{ID: "1", Name: "Ceiling Lightbulb", Funcs: Functionalities{
			Functionality{Description: "Toggle on and off", Path: "/toggle"},
		}},
	}
	fmt.Println("All Devices Endpoint hit!")
	json.NewEncoder(w).Encode(devices)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true).PathPrefix("/api").Subrouter()
	router.HandleFunc("/", handleHome).Methods("GET")
	router.HandleFunc("/devices", allDevices).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	handleRequest()
}
