package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Alert struct {
	Status string `json:"status"`
}

func alertHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var alert Alert
	err = json.Unmarshal(body, &alert)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	log.Printf("Received alert: %s", alert.Status)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Alert received")
}

func main() {
	http.HandleFunc("/webhook", alertHandler)
	port := "8080"
	log.Printf("Starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
