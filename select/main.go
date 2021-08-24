package main

import (
	"fmt"
	"log"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func main() {
	log.Println("Select in Golang")
	output1 := make(chan string)
	output2 := make(chan string)
	ch := make(chan string)

	go server1(output1)
	go server2(output2)

	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}

	go process(ch)

	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case c1 := <-ch:
			fmt.Printf("Received value: %s", c1)
			return
		default:
			fmt.Println("No value received")
		}
	}
	// select{} Select will block until one of its cases is executed.
	// If select does not have any cases it will block forever resulting in a deadlock
	// Always add a default case to select{} to avoid a deadlock
}
