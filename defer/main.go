package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/pashamakhilkumarreddy/golang-snippets/defer/employee"
	"github.com/pashamakhilkumarreddy/golang-snippets/defer/geometry"
)

func finished() {
	fmt.Println("Done printing all names")
}

func printNames(names []string) {
	defer finished()
	for _, name := range names {
		fmt.Printf("Hola %v\n", name)
	}
}

func printA(a int) {
	fmt.Println()
	fmt.Println("value of a in the deffered function is", a)
}

func main() {
	log.Println("Defer in Golang") // mutltiple defer function calls are processed in Last in First out order
	names := []string{"Gwen Stacy", "Ashido Kano", "Amanda"}
	printNames(names)

	emp1 := employee.NewEmployee("Gwen", "Stacy", 1000000, "CEO")

	defer fmt.Println(emp1.FullName())
	fmt.Println("Beinvenidos")

	a := 6
	defer printA(a)
	a = 12
	fmt.Println("value of a before deferred function call", a)

	name := "Benjamin"
	fmt.Printf("Original string: %s\n", string(name))
	fmt.Printf("Reversed string: ")
	//lint:ignore SA6003 ignore for now
	//lint:ignore S1029 ignore for now
	for _, v := range []rune(name) {
		defer fmt.Printf("%c", v)
	}

	var wg sync.WaitGroup
	r1 := geometry.NewRect(12, 15)
	r2 := geometry.NewRect(-6, 10)
	r3 := geometry.NewRect(9, 4)

	rects := geometry.Rects(r1, r2, r3)

	for _, v := range rects {
		wg.Add(1)
		go v.Area(&wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")
}
