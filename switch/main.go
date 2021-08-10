package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	log.Println("Switch in Golang")
	finger := 3
	fmt.Printf("Finger %d is ", finger)
	switch finger { // duplicate cases are not allowed
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

	switch letter := "i"; letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	num := 75
	switch {
	case num >= 0 && num <= 50:
		fmt.Printf("%d is greater than 0 and less than 50\n", num)
	case num >= 51 && num <= 100:
		fmt.Printf("%d is greater than 51 and less than 100\n", num)
	case num >= 101:
		fmt.Printf("%d is greater than 100\n", num)
	}

	switch num := number(); {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is less than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is less than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is less than 200\n", num)
	}

randLoop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Printf("Generated even number %d", i)
			break randLoop
		}
	}
}

func number() (num int) {
	num = 15 * 3
	return
}
