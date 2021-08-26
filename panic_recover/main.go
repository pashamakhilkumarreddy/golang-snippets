package main

import (
	"fmt"
	"log"
	"runtime/debug"
)

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
		debug.PrintStack()
	}
}

func fullName(firstName, lastName *string) {
	defer recovery()
	defer fmt.Println("deferred call in fullName")
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func recoverInValidAccess() {
	if r := recover(); r != nil {
		fmt.Println("recover ", r)
	}
}

func slicePanic() {
	defer recoverInValidAccess()
	n := []int{1, 2, 3}
	fmt.Println(n[3])
	fmt.Println("normally returned from slicePanic")
}

func sum(a, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	<-done
}

func divide(a, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b)
	done <- true
}

func main() {
	log.Println("Panic and Recover in Golang") // recover only works when it is called from the same goroutine which is panicking
	defer fmt.Println("defer called in main")
	firstName := "Gwen"
	fullName(&firstName, nil)
	slicePanic()
	sum(9, 0)
	fmt.Println("returned normally from main")
}
