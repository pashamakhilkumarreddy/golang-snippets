package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Bubble Sort")
	fmt.Println("Please enter the number of elements to enter in the list")
	var length uint64
	_, err := fmt.Scanf("%d", &length)
	if err != nil || length < 0 {
		printMessage("Please enter a valid number")
		os.Exit(1)
	}
	values := genRandomValues(length)
	fmt.Printf("Unsorted list is %v\n", values)
	sortedList := bubbleSort(values)
	fmt.Printf("Sorted list is %v\n", sortedList)
}

func printMessage(message string) {
	fmt.Println(message)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func genUniqueValues(values []int) []int {
	if len(values) <= 1 {
		fmt.Println("List is already unique")
		return make([]int, 0)
	}
	uniqMap := make(map[int]bool)
	list := []int{}
	for _, val := range values {
		if _, value := uniqMap[val]; !value {
			uniqMap[val] = true
			list = append(list, val)
		}
	}
	return list
}

func genRandomValues(length uint64) []int {
	if length <= 0 {
		fmt.Println("Can't fill an empty list")
		return make([]int, 0)
	}
	const MIN = 0
	var MAX = int(length) + 10
	rand.Seed(time.Now().UTC().UnixNano())
	values := make([]int, 0)
	for i := 0; i < MAX; i++ {
		values = append(values, randomInt(MIN, MAX))
	}
	return genUniqueValues(values)
}

func bubbleSort(values []int) []int {
	var arrayLen = len(values)
	if arrayLen <= 1 {
		fmt.Println("Array is alredy sorted")
		return values
	}
	var iterations int = 0
	var swapped bool
	for i := 0; i < arrayLen-1; i++ {
		swapped = false
		for j := 0; j < arrayLen-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapped = true
			}
		}
		if swapped != true {
			break
		}
		iterations++
	}
	fmt.Printf("Completed sorting in %d iterations\n", iterations)
	return values
}
