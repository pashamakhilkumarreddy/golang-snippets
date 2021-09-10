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

	ssns := expressLine.Manifest()
	fmt.Printf("This bus carries %d passengers, here are their SSN's: %v\n", len(ssns), ssns)
}
