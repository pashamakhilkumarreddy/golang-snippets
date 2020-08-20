package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Linear Search")
	var length uint64
	fmt.Print("Please enter the number of values to add:")
	_, err := fmt.Scanf("%d", &length)
	var values []uint64
	if err == nil {
		values = generateNumbers(length)
	} else {
		fmt.Println("Please enter a valid number")
	}
	fmt.Printf("The generated values are %v\n", values)
	var elementToFind uint64
	fmt.Println("Please enter the element to find in the array")
	_, elemErr := fmt.Scanf("%d", &elementToFind)
	if elemErr == nil {
		var result = linearSearch(values, elementToFind)
		fmt.Println(result)
	} else {
		fmt.Println("Please enter a valid number")
	}
}

func generateNumbers(length uint64) []uint64 {
	var (
		MIN uint64 = 0;
		MAX uint64 = length + 100;
	)
	arr := make([]uint64, length)
	for i := 0; uint64(i) < length; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		var randValue = rand.Intn(int(MAX - MIN)) + int(MIN)
		arr[i] = uint64(randValue)
	}
	return arr
}

func linearSearch(values []uint64, elementToFind uint64) string {
	if len(values) == 0 {
		return fmt.Sprintf("Can't find the number in an empty array")
	}
	for i, val := range values {
		if val == elementToFind {
			return fmt.Sprintf("Found %v at position %v in %v steps", elementToFind, i+1, i+1)
		}
	}
	return fmt.Sprintf("Element %v is not in the array", elementToFind)
}
