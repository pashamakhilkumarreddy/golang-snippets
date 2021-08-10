package main

import (
	"fmt"
	"log"
)

func main() {
	log.Println("Loops in Golang")
	for i := 1; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	log.Println()
	for i := 0; i <= 10; i++ {
		if i > 5 {
			break
		}
		fmt.Printf("%d ", i)
	}
	log.Println()
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	log.Println()
	n := 5
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	log.Println()
outer:
	for i := 0; i <= 5; i++ {
		for j := 1; j < 10; j++ {
			if i == j {
				break outer
			}
			fmt.Printf("%d %d ", i, j)
		}
	}
	log.Println()
	num := 0
	for num < 10 {
		fmt.Printf("%d ", num)
		num++
	}
	log.Println()
	k := 0
	for k < 10 {
		fmt.Printf("%d ", k)
		k++
	}
	log.Println()
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
	// for { // Infinite for loop
	// 	fmt.Println("Hola Mundo!")
	// }
}
