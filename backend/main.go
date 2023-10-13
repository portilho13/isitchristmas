package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Message struct  {
	Title string `json:"Title"`
	Ischristmas bool `json:"ischristmas"`
}

type Messages []Message

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Ola")
}

func encodeTeste(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	month := int(currentTime.Month())
	day := currentTime.Day()
	var isx bool
	if month == 12 && day == 25 {
		isx = true
	} else {
		isx = false
	}
	message := Messages{
		Message{Title:"test", Ischristmas:isx},
	}
	json.NewEncoder(w).Encode(message)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/api", encodeTeste)
	log.Fatal(http.ListenAndServe(":8080", nil))
}