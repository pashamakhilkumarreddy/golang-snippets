package main

import (
	"fmt"
	"log"
	"reflect"
)

type employee struct {
	salary  int
	country string
}

func main() {
	log.Println("Maps in Golang") // Maps are reference types,
	// we can't use == to check maps equality (only used to check if the map is nil
	employeeSalary := make(map[string]int)
	fmt.Println(employeeSalary)
	employeeSalary["steve"] = 1200000
	employeeSalary["nick"] = 1500000
	employeeSalary["alice"] = 1300000
	employeeSalary["shaik"] = 12000
	fmt.Println(employeeSalary)

	employeeSalaryMap := map[string]int{
		"stacy":  15000000,
		"ashido": 12000000,
	}
	fmt.Printf("Employee Salary %v\n", employeeSalaryMap)

	var nilMap map[string]int
	fmt.Println(nilMap)
	// nilMap["raizo"] = 12 Cannot add to a nil map

	employeeKey := "stacy"
	salary := employeeSalaryMap[employeeKey]
	fmt.Println(salary)
	fmt.Println(employeeSalaryMap["raizo"])

	val, ok := employeeSalaryMap["ashido"]
	if ok { // ok == true
		fmt.Println(val)
	}

	for key, val := range employeeSalary {
		fmt.Printf("Employee %s's salary is %d\n", key, val)
	}

	delete(employeeSalary, "shaik")
	fmt.Printf("Employee Salary map after deletion %v", employeeSalary)

	emp1 := employee{
		salary:  1200000,
		country: "India",
	}

	emp2 := employee{
		salary:  900000,
		country: "Italy",
	}

	emp3 := employee{
		salary:  1500000,
		country: "Spain",
	}

	employees := map[string]employee{
		"Steve": emp1,
		"Nick":  emp2,
		"Akash": emp3,
	}

	for name, info := range employees {
		fmt.Printf("Employee %s is from %s and earns %d\n", name, info.country, info.salary)
	}

	fmt.Printf("Length of employees map is %v\n", len(employees))

	var emptyMap map[employee]string
	fmt.Println(emptyMap == nil)

	map1 := make(map[string]employee)
	map2 := make(map[employee]string)

	areMapsEqual := reflect.DeepEqual(map1, map2)
	if areMapsEqual {
		fmt.Println("Two maps map1 and map2 are equal")
	} else {
		fmt.Println("Two maps map1 and map2 are not equal")
	}
}
