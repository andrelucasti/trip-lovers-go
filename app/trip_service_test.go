package app

import (
	"github.com/andrelucasti/trip-lovers-go/repository"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var repositoryInMemoryTest = new(repository.TripInMemoryRepository)

func TestNewTrip(t *testing.T) {
	idUser := uuid.New()
	tripRequest := TripRequest{
		Title:     "My First Travel",
		About:     "My Golang travel",
		Departure: "2024-03-01",
		Returns:   "2024-03-07",
	}

	Save(tripRequest, idUser, repositoryInMemoryTest)

	trips := FindAll(repositoryInMemoryTest, idUser)

	expected := TripResponse{
		Title:     "My First Travel",
		About:     "My Golang travel",
		Departure: "2024-03-01",
		Returns:   "2024-03-07",
	}

	assert.Contains(t, trips, expected)
}

func TestTripsByUserId(t *testing.T) {
	idUser := uuid.New()
	tripRequest := TripRequest{
		Title:     "User's travel",
		About:     "User's travel about",
		Departure: "2024-03-01",
		Returns:   "2024-03-07",
	}

	Save(tripRequest, idUser, repositoryInMemoryTest)

	idUser2 := uuid.New()
	tripRequest2 := TripRequest{
		Title:     "User's 2 travel",
		About:     "User's 2 travel about",
		Departure: "2023-05-01",
		Returns:   "2023-10-07",
	}

	Save(tripRequest2, idUser2, repositoryInMemoryTest)

	trips := FindAll(repositoryInMemoryTest, idUser)

	expected := TripResponse{
		Title:     "User's travel",
		About:     "User's travel about",
		Departure: "2024-03-01",
		Returns:   "2024-03-07",
	}

	assert.Equal(t, 1, len(trips))
	assert.Contains(t, trips, expected)
}
