package app

import (
	"github.com/andrelucasti/trip-lovers-go/business"
	"time"
)

func Save(tripRequest TripRequest, repository business.TripRepository) {
	departure, _ := time.Parse(time.DateOnly, tripRequest.departure)
	returns, _ := time.Parse(time.DateOnly, tripRequest.returns)

	newTrip := business.NewTrip(tripRequest.title, tripRequest.about, departure, returns)

	repository.Save(newTrip)
}

func FindAll(repository business.TripRepository) []TripRequest {
	var result []TripRequest

	trips := repository.FindAll()

	for i := 0; i < len(trips); i++ {
		trip := trips[i]

		result = append(result, TripRequest{
			trip.Title,
			trip.About,
			trip.Departure.Format(time.DateOnly),
			trip.Returns.Format(time.DateOnly),
		})
	}

	return result
}
