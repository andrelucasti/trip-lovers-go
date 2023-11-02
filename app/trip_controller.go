package app

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewTrip(w http.ResponseWriter, r *http.Request) {
	var tripRequest TripRequest

	err := json.NewDecoder(r.Body).Decode(&tripRequest)

	if err != nil {
		log.Fatalf("Got an error at create a new trip - %s", err.Error())
	}
}
