package main

import (
	"fmt"
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/oops/services"
)

func main() {
	log.Println("Oops in Golang")
	expressLine := services.Bus{
		Name: "Express Line",
	}
	expressLine.Passengers = make([]services.Passenger, 0)
	expressLine.Passengers = append(expressLine.Passengers, services.Passenger{SSN: "001"})
	expressLine.Passengers = append(expressLine.Passengers, services.Passenger{SSN: "002"})
	expressLine.Passengers = append(expressLine.Passengers, services.Passenger{SSN: "003"})

	ssns := make([]string, 0)
	for _, p := range expressLine.Passengers {
		ssns = append(ssns, p.SSN)
	}
	fmt.Printf("The bus carries %d passengers, here are their SSN's: %v\n", len(ssns), ssns)
}
