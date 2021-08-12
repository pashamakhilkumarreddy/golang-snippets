package main

import (
	"fmt"
	"log"
)

func printArray(arr [3]int) {
	arrNew := arr
	arrNew[2] = 20
	fmt.Println(arr, arrNew)
}

func printArr(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// cannot assign arrays of different lengths to each other a := [5]int{1, 2, 3, 4, 5}; var b [25]int, can't do a = b
	// arrays are passed by value in Golang
	log.Println("Arrays and Slices in Golang")
	var a [3]int
	fmt.Println(a)
	a[0], a[1], a[2] = 1, 2, 3
	fmt.Println(a)

	b := [3]int{3, 6, 9}
	fmt.Println(b)

	c := [3]int{12}
	fmt.Println(c)

	d := [...]int{12, 24, 36}
	fmt.Println(d)

	printArray(d)

	e := [...]float64{63.3, 36.9, 99.1}
	fmt.Println(e, len(e))

	for i := 0; i < len(e); i++ {
		fmt.Printf("%d th element of a is %.2f \n", i, e[i])
	}

	f := [...]float64{12.6, 19.6, 21.3, 36.9, 42.6, 54.9}
	sum := float64(0)
	for i, v := range f {
		fmt.Printf("%d the element of f is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of f", sum)

	g := [3][2]string{
		{"lion", "tiger"},
		{"cat", "dog"},
		{"pigeon", "peacock"},
	}

	printArr(g)

	var h [3][2]string
	h[0][0] = "apple"
	h[0][1] = "facebook"
	h[1][0] = "microsoft"
	h[1][1] = "google"
	h[2][0] = "amazon"
	h[2][1] = "intel"
	printArr(h)

	i := [6]int{1, 2, 3, 4, 5, 6}
	var iSlice []int = i[2:5]
	fmt.Println(iSlice, len(iSlice), cap(iSlice))

	j := []int{1, 2, 3}
	fmt.Println(j)

	k := make([]int, 5, 10)
	fmt.Println(k)
	k = append(k, j...)
	fmt.Println(k)

	countries := []string{"USA", "Germany", "France", "Spain", "Italy", "UK"}
	fmt.Println(countries)
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(countries))
	copy(countriesCpy, neededCountries)
	fmt.Println(countriesCpy)
}
