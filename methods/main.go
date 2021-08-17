package main

import (
	"fmt"
	"log"
	"math"
)

type employee struct {
	name     string
	salary   int
	currency string
}

func (e employee) displaySalary() {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

func (e employee) changeName(name string) {
	//lint:ignore SA4005 demo-purpose
	e.name = name
}

func (e *employee) changeSalary(salary int) {
	e.salary = salary
}

func displaySalary(e *employee) {
	fmt.Printf("Salary of %s is %s%d\n", e.name, e.currency, e.salary)
}

type rectangle struct {
	length, width float64
}

type circle struct {
	radius float64
}

func area(r rectangle) float64 {
	return r.length * r.width
}

func (r rectangle) Area() float64 {
	return r.length * r.width
}

func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

type address struct {
	city, state string
}

type person struct {
	firstName, lastName string
	address
}

func (a address) fullAddress() {
	fmt.Printf("Full address %s %s\n", a.city, a.state)
}

type myInt int

// Method on non-struct types
// Receiver type and definition should be present in the same package
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	log.Println("Methods in Golang")
	emp1 := employee{
		name:     "Gwen Stacy",
		salary:   200000,
		currency: "$",
	}
	emp1.displaySalary()
	displaySalary(&emp1)

	fmt.Println("Employee before change", emp1)

	emp1.changeName("Ashido")
	emp1.changeSalary(300000)
	fmt.Println("Employee after change", emp1)

	rect := rectangle{10, 20}
	fmt.Println("Rectangle area is", rect.Area())

	circ := circle{
		radius: 6,
	}
	fmt.Println("Circle area is", circ.Area())

	p1 := person{
		firstName: "Ashido",
		lastName:  "Kano",
		address: address{
			city:  "Dallas",
			state: "Texas",
		},
	}
	p1.fullAddress() // p1.address.fullAddress()

	r := &rect

	fmt.Println("Rectangle area (with function) is", area(rect))
	// fmt.Println("Rectangle area (with function) is", area(r)) : Not possible

	fmt.Println("Rectangle area is", rect.Area())
	fmt.Println("Rectangle area is", r.Area())

	a := myInt(6)
	b := myInt(9)

	fmt.Println("Sum is", a.add(b))
}
