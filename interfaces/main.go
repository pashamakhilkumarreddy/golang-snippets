package main

import (
	"fmt"
	"log"
)

type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// MyString implements VowelsFinder
func (str MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range str {
		if rune == 'a' || rune == 'b' || rune == 'c' || rune == 'd' || rune == 'e' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

type SalaryCalculator interface {
	CalculateSalary() int
}

type permanent struct {
	empId    int
	basicpay int
	pf       int
}

type contract struct {
	empId    int
	basicpay int
}

type freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

// salary of a permanent employee is the sum of basic pay and pf
func (p permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of a contract employee is the basic pay alone
func (c contract) CalculateSalary() int {
	return c.basicpay
}

// salary of a freelancer is no of hours worked * rate per hour
func (f freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

// Total Expenses of all employees
func totalExpenses(s []SalaryCalculator) {
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total expense per month is %d$\n", expense)
}

type worker interface {
	Work()
}

type person struct {
	name string
}

func (p person) Work() {
	fmt.Println(p.name, "is working")
}

func describe(w worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func desc(i interface{}) {
	fmt.Printf("Type = %T, value = %v\n", i, i)
}

func assert(i interface{}) {
	s := i.(int) // to get the underlying value from i
	fmt.Println(s)
}

func assertStr(i interface{}) {
	v, ok := i.(string)
	fmt.Println(v, ok)
}

func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i)
	case int:
		fmt.Printf("I am an int and my value is %d\n", i)
	default:
		fmt.Println("Unknown type")
	}
}

type describer interface {
	describe()
}

type persona struct {
	name string
	age  int
}

func (p persona) describe() {
	fmt.Printf("%s is %d years old", p.name, p.age)
}

func findTypeI(i interface{}) {
	switch v := i.(type) {
	case describer:
		v.describe()
	default:
		fmt.Println("Unknown type")
	}
}

func main() {
	log.Println("Interfaces in Golang")
	name := MyString("Ashido Kano")
	// v := name
	fmt.Printf("Vowels are %c\n", name.FindVowels())

	pemp1 := permanent{
		empId:    1,
		basicpay: 3000,
		pf:       15,
	}

	pemp2 := permanent{
		empId:    2,
		basicpay: 6000,
		pf:       45,
	}

	pemp3 := permanent{
		empId:    3,
		basicpay: 4500,
		pf:       30,
	}

	cemp1 := contract{
		empId:    4,
		basicpay: 2000,
	}

	femp1 := freelancer{
		empId:       5,
		ratePerHour: 24,
		totalHours:  150,
	}

	femp2 := freelancer{
		empId:       6,
		ratePerHour: 30,
		totalHours:  300,
	}

	employees := []SalaryCalculator{
		pemp1,
		pemp2,
		pemp3,
		cemp1,
		femp1,
		femp2,
	}

	totalExpenses(employees)

	p := person{
		name: "Raizo",
	}
	var w worker = p
	describe(w)
	w.Work()

	s := "Hola Mundo"
	desc(s)
	i := 33
	desc(i)
	strt := struct {
		name string
	}{
		name: "Amanda",
	}
	desc(strt)

	var inte interface{} = 99
	assert(inte)

	var a interface{} = "namen"

	assertStr(a)

	findType("Alexis")
	findType(333)
	findType(3.3)

	per := persona{
		name: "Lorenzo",
		age:  42,
	}
	findTypeI(per)
}
