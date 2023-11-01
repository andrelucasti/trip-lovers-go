package app

import (
	"github.com/andrelucasti/trip-lovers-go/business"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrip(t *testing.T) {
	tripRequest := TripRequest{
		title:     "My First Travel",
		about:     "My Golang travel",
		departure: "2024-03-01",
		returns:   "2024-03-07",
	}

	repository := new(business.TripInMemoryRepository)

	Save(tripRequest, repository)

	trips := FindAll(repository)

	assert.Contains(t, trips, tripRequest)
}
