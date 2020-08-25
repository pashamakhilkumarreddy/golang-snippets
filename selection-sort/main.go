package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	fmt.Println("Selection sort")
	fmt.Println("Please enter the number of elements to enter in the list")
	var length uint64
	_, err := fmt.Scanf("%d", &length)
	if err != nil || length < 0 {
		printMessage("Please enter a valid number")
		os.Exit(1)
	}
	values := generateRandomValues(int(length))
	fmt.Printf("Unsorted list is %v\n", values)
	sortedList := selectionSort(values)
	fmt.Printf("Sorted List is %v\n", sortedList)
}

func printMessage(message string) {
	fmt.Println(message)
}

func randomInt(min, max int) int {
	return min + rand.Intn(max - min)
}

func genUniqValues(values []int) []int {
	if len(values) <= 1 {
		fmt.Println("List is already unique")
		return values
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

func generateRandomValues(length int) []int {
	if length <= 0 {
		fmt.Println("Can't generate an empty list")
		return make([]int, 0)
	}
	var (
		MIN = 0
		MAX = length + 20
	)
	values := make([]int, 0)
	rand.Seed(time.Now().UTC().UnixNano())
	for i:=0; i< length ;i++ {
		values = append(values, randomInt(MIN, MAX))
	}
	uniqValues := genUniqValues(values)
	return uniqValues
}

func selectionSort(values []int) []int {
	var length = len(values)
	if length <= 1 {
		fmt.Println("The list is already sorted")
		return values
	}
	for i := 0; i < length - 1; i++ {
		minIndex := i
		for j := i + 1 ; j < length; j++ {
			if values[j] <= values[minIndex] {
				minIndex = j
			}
		}
		if (minIndex != i) {
			var temp = values[minIndex]
			values[minIndex] = values[i]
			values[i] = temp
		}
	}
	return values
}