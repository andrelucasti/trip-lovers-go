package app

type TripRequest struct {
	Title     string `json:"title"`
	About     string `json:"about"`
	Departure string `json:"departure"`
	Returns   string `json:"returns"`
}
