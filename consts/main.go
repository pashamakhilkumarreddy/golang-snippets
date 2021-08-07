package main

import (
	"fmt"
	_ "math"
)

func main() {
	fmt.Println("Constants in Golang")
	const a = 500
	fmt.Println(a)
	const (
		name    = "Gwen Stacy"
		age     = 27
		country = "Neverland"
		general = true
	)
	// const sqrt = math.Sqrt(25) // Not allowed as this call happens during run time
	fmt.Println(name, age, country, general)
	var fa float64 = a
	var fc complex128 = a
	fmt.Println(fa, fc)
}
