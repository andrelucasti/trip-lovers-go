package business

var data []Trip

type TripRepository interface {
	Save(trip Trip)
	FindAll() []Trip
}

type TripInMemoryRepository struct {
}

func (tr TripInMemoryRepository) Save(trip Trip) {
	data = append(data, trip)
}

func (tr TripInMemoryRepository) FindAll() []Trip {
	return data
}
