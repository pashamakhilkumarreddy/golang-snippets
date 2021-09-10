package services

import "fmt"

// Passenger represents a bus passenger, uniquely identified by their SSN
type Passenger struct {
	SSN        string
	SeatNumber uint8
}

// Bus carriers Passengers from A to B if they have a valid bus ticket
type Bus struct {
	name       string
	passengers Passengers
}

type Passengers map[string]Passenger

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

func NewPassengerSet() Passengers {
	return make(map[string]Passenger)
}

func (p Passengers) Visit(visitor func(Passenger)) {
	for _, one := range p {
		visitor(one)
	}
}

func (p Passengers) Find(ssn string) Passenger {
	if v, ok := p[ssn]; ok {
		return v
	}
	return Passenger{}
}

func (p *Passengers) VisitUpdate(visitor func(*Passenger)) {
	for ssn, pp := range *p {
		visitor(&pp)
		(*p)[ssn] = pp
	}
}

func (p Passengers) Manifest() []string {
	ssns := make([]string, 0, len(p))
	p.Visit(func(p Passenger) { ssns = append(ssns, p.SSN) })
	return ssns
}

func NewBus(name string) Bus {
	b := Bus{}
	b.name = name
	b.passengers = NewPassengerSet()
	return b
}

func (b *Bus) Add(p Passenger) {
	if b.passengers == nil {
		b.passengers = make(map[string]Passenger)
	}
	b.passengers[p.SSN] = p
	fmt.Printf("%s: boarded passenger with SSN%q\n", b.name, p.SSN)
}

func (b *Bus) Remove(p Passenger) {
	delete(b.passengers, p.SSN)
	fmt.Printf("%s: unboarded passenger with SSN %q\n", b.name, p.SSN)
}

func (b Bus) Manifest() []string {
	return b.passengers.Manifest()
}

func (b *Bus) VisitPassengers(visitor func(Passenger)) {
	b.passengers.Visit(visitor)
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
