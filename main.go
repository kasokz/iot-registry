package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Device struct {
	ID   string
	Name string
	Kind string
}

type Devices []Device

func allDevices(w http.ResponseWriter, r *http.Request) {
	devices := Devices{
		Device{ID: "1", Name: "Ceiling Lightbulb", Kind: "LightBulb"},
	}
	fmt.Println("All Devices Endpoint hit!")
	json.NewEncoder(w).Encode(devices)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func handleRequest() {
	http.HandleFunc("/api", handleHome)
	http.HandleFunc("/devices", allDevices)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequest()
}
