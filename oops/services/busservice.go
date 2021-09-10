package services

// Passenger represents a bus passenger, uniquely identified by their SSN
type Passenger struct {
	SSN        string
	SeatNumber uint8
}

// Bus carriers Passengers from A to B if they have a valid bus ticket
type Bus struct {
	name       string
	passengers map[string]Passenger
}

func NewPassenger(SSN string, SeatNumber interface{}) Passenger {
	p := Passenger{}
	p.SSN = SSN
	switch v := SeatNumber.(type) {
	case int:
	case uint8:
		p.SeatNumber = uint8(v)
	}
	return p
}

func NewBus(name string) Bus {
	b := Bus{}
	b.name = name
	return b
}

func (b *Bus) Add(p Passenger) {
	if b.passengers == nil {
		b.passengers = make(map[string]Passenger)
	}
	b.passengers[p.SSN] = p
}

func (b *Bus) VisitPassengers(visitor func(Passenger)) {
	for _, p := range b.passengers {
		visitor(p)
	}
}

func (b *Bus) FindPassenger(ssn string) Passenger {
	if v, ok := b.passengers[ssn]; ok {
		return v
	}
	return Passenger{}
}

func (b *Bus) UpdatePassenger(visitor func(*Passenger)) {
	ps := make(map[string]Passenger, len(b.passengers))
	for ssn, p := range b.passengers {
		visitor(&p)
		ps[ssn] = p
	}
	b.passengers = ps
}
