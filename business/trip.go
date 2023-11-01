package business

import (
	"github.com/google/uuid"
	"time"
)

type Trip struct {
	Id        uuid.UUID
	Title     string
	About     string
	Departure time.Time
	Returns   time.Time
}

func NewTrip(title string, about string, departure time.Time, returns time.Time) Trip {
	return Trip{
		uuid.New(),
		title,
		about,
		departure,
		returns,
	}
}
