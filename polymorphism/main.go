package main

import (
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/polymorphism/project"
)

func main() {
	log.Println("Polymorphism in Golang")
	project1 := project.NewProject("Ultra Instinct", 10000)
	project2 := project.NewProject("Super Saiyan", 8000)
	project3 := project.NewFreeLance("Normalmente", 600, 10)
	project4 := project.NewDigitalRevenue("You Tube", 60000, 1.2)

	incomeStreams := []project.Income{project1, project2, project3, project4}
	project.CalculateNetIncome(incomeStreams)
}
