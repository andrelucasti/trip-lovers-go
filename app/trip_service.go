package app

import (
	"github.com/andrelucasti/trip-lovers-go/business"
	"time"
)

func Save(tripRequest TripRequest, repository business.TripRepository) {
	departure, _ := time.Parse(time.DateOnly, tripRequest.Departure)
	returns, _ := time.Parse(time.DateOnly, tripRequest.Returns)

	newTrip := business.NewTrip(tripRequest.Title, tripRequest.About, departure, returns)

	repository.Save(newTrip)
}

func FindAll(repository business.TripRepository) []TripResponse {
	var result []TripResponse

	trips := repository.FindAll()

	for i := 0; i < len(trips); i++ {
		trip := trips[i]

		result = append(result, TripResponse{
			trip.Title,
			trip.About,
			trip.Departure.Format(time.DateOnly),
			trip.Returns.Format(time.DateOnly),
		})
	}

	return result
}
