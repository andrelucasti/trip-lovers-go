package app

import (
	"github.com/andrelucasti/trip-lovers-go/repository"
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

	repositoryInMemory := new(repository.TripInMemoryRepository)

	Save(tripRequest, repositoryInMemory)

	trips := FindAll(repositoryInMemory)

	expected := TripResponse{
		title:     "My First Travel",
		about:     "My Golang travel",
		departure: "2024-03-01",
		returns:   "2024-03-07",
	}
	assert.Contains(t, trips, expected)
}
