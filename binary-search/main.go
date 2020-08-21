package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Binary Search")
	fmt.Println("Please enter the number of values to enter in the array")
	var length int
	_, err := fmt.Scanf("%d", &length)
	var values []int
	if err != nil {
		printMessage(err, "Please enter a valid number")
	}
	values = generateList(length)
	fmt.Printf("Value are %v\n", values)
	fmt.Println("Please enter the value to search in the list")
	var elementToFind int
	_, elementErr := fmt.Scanf("%d", &elementToFind)
	if elementErr != nil {
		printMessage(elementErr, "Please enter a valid number")
	}
	var result string = binarySearch(values, elementToFind)
	fmt.Println(result)
}

func printMessage(err interface{}, message string) {
	if err != nil {
		fmt.Println(message)
		os.Exit(1)
	}
}

func generateList(length int) []int {
	if length == 0 {
		fmt.Println("Can't generate an empty list")
		return nil
	}
	arr := make([]int, 0)
	for i := 0; i < length; i++ {
		arr = append(arr, i)
	}
	return arr
}

func binarySearch(values []int, element int) string {
	var (
		low        int = 0
		high       int = len(values) - 1
		comparison int = 0
	)
	for low <= high {
		mid := low + (high-low)/2
		comparison++
		midValue := values[mid]
		if element == midValue {
			return fmt.Sprintf("Found %d at position %d in %d comparisons",
				element, mid + 1, comparison)
		} else if element > midValue {
			low = mid + 1
		} else if element < midValue {
			high = mid - 1
		}
	}
	return fmt.Sprintf("Couldn't find the element %d in the list", element)
}
