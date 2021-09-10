package main

import (
	"fmt"
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/oops/services"
)

func main() {
	log.Println("Oops in Golang")
	expressLine := services.NewBus("Express Line")
	expressLine.Add(services.NewPassenger("001", nil))
	expressLine.Add(services.NewPassenger("002", nil))
	expressLine.Add(services.NewPassenger("003", nil))

	ssns := make([]string, 0)
	expressLine.VisitPassengers(func(p services.Passenger) { ssns = append(ssns, p.SSN) })
	fmt.Printf("The bus carries %d passengers, here are their SSN's: %v\n", len(ssns), ssns)

	var seatNumber uint8
	expressLine.UpdatePassenger(func(p *services.Passenger) {
		seatNumber++
		p.SeatNumber = seatNumber
	})

	expressLine.VisitPassengers(func(p services.Passenger) {
		fmt.Printf("Passenger with SSN %q has seat %d\n", p.SSN, p.SeatNumber)
	})
}
