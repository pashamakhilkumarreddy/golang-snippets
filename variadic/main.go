package main

import (
	"fmt"
	"log"
)

func find(num int, nums ...int) bool {
	fmt.Printf("The number to find is %d\n", num)
	if len(nums) == 0 {
		fmt.Println("Nothing to search in")
		return false
	}
	found := false
	for i, v := range nums {
		if v == num {
			found = true
			fmt.Println(num, "found at index", i, "in", nums)
			break
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	return found
}

func main() {
	log.Println("Variadic functions in Golang")
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	_ = find(6, nums...)
	_ = find(9, 2, 5, 6, 10, 33, 66, 92, 75, 100, 12)
	_ = find(33)
	_ = find(5, []int{}...)
}
