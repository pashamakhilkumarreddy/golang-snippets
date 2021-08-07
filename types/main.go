package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Types in Golang")
	a := true
	b := false
	fmt.Println("a: ", a, " b : ", b)
	c := a && b
	fmt.Println("c :", c)
	d := a || b
	fmt.Println("d: ", d)
	var e int = 65
	f := 95
	fmt.Println("value of e is ", e, " value of f is ", f)
	fmt.Printf("Type of a is %T \n", a)
	fmt.Printf("Type of b is %T and size of b is %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("Type of f is %T and size of f is %d\n", f, unsafe.Sizeof(f))
	var g float64 = 54.12
	h := 12.3
	sum := g + h
	diff := g - h
	fmt.Println("Sum of g + h is ", sum, " Difference of g - h is ", diff)
	i := 9 + 30i
	j := complex(3, 12)
	cadd := i - j
	fmt.Println("Sum ", cadd)
	cmul := i * j
	fmt.Println("Mul ", cmul)
	first := "Gwen"
	last := "Stacy"
	name := first + " " + last
	fmt.Println("Name is ", name)
	fmt.Println("Addition with casting ", e+int(h))
}
