package repository

import (
	"github.com/andrelucasti/trip-lovers-go/business"
	"log"
)

var data []business.Trip

type TripInMemoryRepository struct {
}

func (tr TripInMemoryRepository) Save(trip business.Trip) {
	log.Printf("Saving new Trip - %s", trip.Id)
	data = append(data, trip)
}

func (tr TripInMemoryRepository) FindAll() []business.Trip {

	return data
}
