package app

import (
	"encoding/json"
	"github.com/andrelucasti/trip-lovers-go/repository"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

var repositoryInMemory = new(repository.TripInMemoryRepository)

func NewTrip(w http.ResponseWriter, r *http.Request) {
	userIdParam := mux.Vars(r)["userId"]
	userId, errUuid := uuid.Parse(userIdParam)

	if errUuid != nil {
		http.Error(w, "Got an error at parse userId for a new trip", http.StatusBadRequest)
	} else {
		var tripRequest TripRequest
		err := json.NewDecoder(r.Body).Decode(&tripRequest)
		Save(tripRequest, userId, repositoryInMemory)

		if err != nil {
			http.Error(w, "Got an error at create a new trip", http.StatusInternalServerError)
		}

	}

	w.WriteHeader(http.StatusAccepted)
}

func ListAllTripsByUserId(w http.ResponseWriter, r *http.Request) {
	userIdParam := mux.Vars(r)["userId"]
	userId, errUuid := uuid.Parse(userIdParam)

	if errUuid != nil {
		http.Error(w, "Got an error at parse userId for a new trip", http.StatusBadRequest)
	} else {
		tripResponses := FindAll(repositoryInMemory, userId)

		err := json.NewEncoder(w).Encode(tripResponses)

		if err != nil {
			http.Error(w, "Got an error at show the trips from this userId", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}
