package simpleinterest

import "fmt"

func init() {
	fmt.Println("Simple Interest package is initialized")
}

func Calculate(p, r, t float64) (interest float64) {
	interest = p * r * t / 100
	return
}
