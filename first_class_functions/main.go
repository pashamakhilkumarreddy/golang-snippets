package main

import (
	"fmt"
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/first_class_functions/employee"
)

type add func(a, b int) int

func simple(a func(a, b int) int) {
	fmt.Println(a(30, 90))
}

func closureSum() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

func addNum(num int) func(int) int {
	counter := num
	f := func(a int) int {
		counter += a
		return counter
	}
	return f
}

func filter(nums []int, f func(int) bool) []int {
	var res []int
	for _, v := range nums {
		if r := f(v); r {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	log.Println("First Class Functions in Golang")
	func1 := func() {
		fmt.Println("Hola Mundo")
	}
	func1()
	fmt.Printf("Type of func1 is %T\n", func1)
	func() {
		fmt.Println("self invoking function")
	}()

	func(name string) {
		fmt.Printf("Hola! Estoy %s\n", name)
	}("Gwen Stacy")

	var sum add = func(a, b int) int {
		return a + b
	}
	res := sum(3, 6)
	fmt.Println(res)

	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	s := closureSum()
	fmt.Println(s(9, 81))

	namen := "Ashido Kano"
	func() {
		fmt.Printf("My name is %v from closure\n", namen)
	}()

	a := addNum(10)

	b := addNum(20)
	fmt.Println(a(1))
	fmt.Println(a(1))
	fmt.Println(a(1))

	fmt.Println(b(1))
	fmt.Println(b(1))
	fmt.Println(b(1))

	e1 := employee.NewEmployee("Gwen Stacy", 9999999, 25, "CEO", "UK")
	e2 := employee.NewEmployee("Ashido Kano", 450000, 27, "CTO", "USA")
	e3 := employee.NewEmployee("Raizo", 300000, 30, "CFO", "CN")

	f2 := employee.Filter(func(e employee.Employee) bool {
		return e.Salary > 400000
	}, e1, e2, e3)
	fmt.Println(f2)

	filteredNums := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(num int) bool {
		return num%2 == 0
	})

	fmt.Println(filteredNums)
}
