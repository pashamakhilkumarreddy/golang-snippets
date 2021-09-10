package services

// Passenger represents a bus passenger, uniquely identified by their SSN
type Passenger struct {
	SSN        string
	SeatNumber uint8
}

// Bus carriers Passengers from A to B if they have a valid bus ticket
type Bus struct {
	Name       string
	Passengers []Passenger
}
