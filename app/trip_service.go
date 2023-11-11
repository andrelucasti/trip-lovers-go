package app

import (
	"github.com/andrelucasti/trip-lovers-go/business"
	"github.com/google/uuid"
	"time"
)

func Save(tripRequest TripRequest, idUser uuid.UUID, repository business.TripRepository) {
	departure, _ := time.Parse(time.DateOnly, tripRequest.Departure)
	returns, _ := time.Parse(time.DateOnly, tripRequest.Returns)

	newTrip := business.NewTrip(tripRequest.Title, tripRequest.About, departure, returns, idUser)

	repository.Save(newTrip)
}

func FindAll(repository business.TripRepository, idUser uuid.UUID) []TripResponse {
	var result []TripResponse

	trips := repository.FindAll()

	for i := 0; i < len(trips); i++ {
		trip := trips[i]

		if idUser == trip.IdUser {
			result = append(result, TripResponse{
				trip.Title,
				trip.About,
				trip.Departure.Format(time.DateOnly),
				trip.Returns.Format(time.DateOnly),
			})
		}
	}

	return result
}
