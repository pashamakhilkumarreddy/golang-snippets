package main

import (
	"log"
)

func main() {
	log.Println("if-else in Golang")

	if num := 9; num%2 == 0 {
		log.Println("The number ", num, "is even")
	} else {
		log.Println("The number ", num, "is odd")
	}

	num := 12
	if num%2 == 0 {
		log.Println("The number ", num, "is even")
		return
	}
	log.Println("The number ", num, "is odd")

}
