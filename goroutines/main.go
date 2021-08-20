package main

import (
	"fmt"
	"log"
	"time"
)

func hola() {
	fmt.Println("Hola Mundo!")
}

func printNumbers() {
	for i := 0; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}

func printAlphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", i)
	}
}

func main() {
	go hola()
	go printNumbers()
	go printAlphabets()
	time.Sleep(1 * time.Second)
	log.Println("Goroutines in Golang")
	fmt.Println("Main terminated")

}
