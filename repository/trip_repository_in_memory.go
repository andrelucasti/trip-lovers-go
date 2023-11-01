package repository

import "github.com/andrelucasti/trip-lovers-go/business"

var data []business.Trip

type TripInMemoryRepository struct {
}

func (tr TripInMemoryRepository) Save(trip business.Trip) {
	data = append(data, trip)
}

func (tr TripInMemoryRepository) FindAll() []business.Trip {

	return data
}
