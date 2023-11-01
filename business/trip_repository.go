package business

type TripRepository interface {
	Save(trip Trip)
	FindAll() []Trip
}
