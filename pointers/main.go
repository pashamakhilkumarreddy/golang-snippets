package main

import (
	"fmt"
	"log"
)

func change(val *int) {
	*val = 66
}

func hola() *int {
	a := 9
	return &a
}

func modify(arr *[6]int) {
	(*arr)[0] = 9 // arr[0] = 9
}

func modifySlice(arr []int) {
	arr[3] = 33
}

func main() {
	log.Println("Pointers in Golang") // Go doesn't support pointer arithmetic
	b := 333
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Address of b is %v\n", a)
	var c *int
	if c == nil {
		fmt.Println("c is ", c)
		c = &b
		fmt.Println("c after initialization is ", c)
	}

	height := new(float64)
	fmt.Printf("Height value is %f, type is %T, address is %v\n", *height, height, height)
	*height = 180.3
	fmt.Println("New height value is ", *height)

	d := 111
	e := &d
	fmt.Println("Address of d is ", e)
	fmt.Println("Value of d is ", *e)
	*e += 111
	fmt.Println("Value of d is ", *e)

	f := 666
	g := &f
	fmt.Println("Address of f is ", g)
	fmt.Println("Value of f is ", *g)

	change(g)
	fmt.Println("Value of g after f is ", *g)

	h := hola()
	fmt.Println("Value of d", h)

	arr := [6]int{1, 2, 3, 4, 5, 6}
	modify(&arr)
	fmt.Println("Modify as a pointer to an array", arr)
	modifySlice(arr[:])
	fmt.Println("Modify in slice ", arr)
}
