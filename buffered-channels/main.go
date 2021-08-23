package main

import (
	"fmt"
	"log"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to ch")
	}
	close(ch)
}

func main() {
	log.Println("Buffered Channels in Golang")

	ch := make(chan string, 3)

	ch <- "Gwen Stacy"
	ch <- "Ashido Kano"
	ch <- "Raizo"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	ch2 := make(chan int, 2)
	go write(ch2)
	time.Sleep(2 * time.Second)
	for v := range ch2 {
		fmt.Println("read value", v, "from ch")
		time.Sleep(2 * time.Second)
	}
	fmt.Println("capacity of ch channel is ", cap(ch))
	fmt.Println("length is ", len(ch))
}
