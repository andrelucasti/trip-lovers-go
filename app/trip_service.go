package app

var data []TripRequest

func Save(tripRequest TripRequest) {
	data = append(data, tripRequest)
}

func FindAll() []TripRequest {
	return data
}
