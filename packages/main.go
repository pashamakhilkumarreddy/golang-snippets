package main

import (
	"fmt"
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/packages/simpleinterest"
)

var p, t, r = 5000.0, 10.0, 1.0

func init() {
	fmt.Println("Main package is initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	fmt.Println("Packages in Golang")

	si := simpleinterest.Calculate(p, t, r)
	fmt.Println("Simple interest is ", si)
}
