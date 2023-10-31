package app

import (
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

	Save(tripRequest)

	trips := FindAll()

	assert.Contains(t, trips, tripRequest)
}
