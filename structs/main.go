package main

import (
	"fmt"
	"log"

	"github.com/pashamakhilkumarreddy/golang-snippets/structs/models"
)

type employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

type emp struct {
	firstName, lastName string
	age, salary         int
}

// Anonymous structs
type person struct {
	string
	int
}

type address struct {
	city, state string
}

type persona struct {
	name string
	age  int
	addr address
}

type humano struct {
	name    string
	address // anonymous structs are promoted
}

func main() {
	log.Println("Structs in Golang")
	emp1 := employee{
		firstName: "Gwen",
		age:       27,
		salary:    1500000,
		lastName:  "Stacy",
	}
	emp2 := emp{"Ashido", "Kano", 26, 1200000}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Justin",
		age:       30,
		salary:    9000000,
		lastName:  "Paul",
	}
	fmt.Println("Employee 3", emp3)

	fmt.Println("First Name:", emp3.firstName)
	fmt.Println("Last Name:", emp3.lastName)
	fmt.Println("Age:", emp3.age)
	fmt.Println("Salary:", emp3.salary)
	emp3.salary = 14000000
	fmt.Printf("Salary: %d\n", emp3.salary)

	var emp4 emp // zero valued struct
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)

	emp5 := emp{
		firstName: "Raizo",
	}
	fmt.Println("Employee 4:", emp5)

	emp6 := &employee{
		firstName: "Amanda",
		age:       21,
	}
	fmt.Println("Employee 6: ", *emp6, emp6.firstName, (*emp6).firstName)

	p1 := person{
		string: "Teresa",
		int:    42,
	}
	fmt.Println(p1)

	p2 := persona{
		name: "Lorenzo",
		age:  36,
		addr: address{
			city:  "11th Calle",
			state: "Barcelona",
		},
	}

	fmt.Println(p2, p2.addr)

	h1 := humano{
		name: "Marina",
		address: address{
			city:  "Dallas",
			state: "Texas",
		},
	}
	fmt.Println(h1, h1.state)

	s1 := models.Spec{
		Maker: "apple",
		Price: 200000,
	}

	fmt.Println(s1)

	n1 := models.Name{
		FirstName: "Gwen",
		LastName:  "Stacy",
	}

	n2 := models.Name{
		FirstName: "Gwen",
		LastName:  "Stacy",
	}
	// Struct variables are not comparable if they contain fields that are not comparable
	if n1 == n2 {
		fmt.Println("Structs", n1, n2, "are equal")
	} else {
		fmt.Println("Structs", n1, n2, "are not equal")
	}
}
