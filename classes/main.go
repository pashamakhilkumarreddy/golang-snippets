package main

import (
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/classes/employee"
)

func main() {
	log.Println("Structs instead of classes in Golang")
	// var emp Employee
	// emp.LeavesRemaining()
	// e := employee.Employee{"Gwen", "Stacy", 24, 6}
	// e.LeavesRemaining()

	e := employee.NewEmployee("Gwen", "Stacy", 24, 6)
	e.LeavesRemaining()
}
