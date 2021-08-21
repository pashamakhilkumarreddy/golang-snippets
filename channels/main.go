package main

import (
	"fmt"
	"log"
)

func hola(done chan bool) {
	fmt.Println("Hola, in sleep mode")
	// time.Sleep(3 * time.Second)
	fmt.Println("Hola, in active mode")
	fmt.Println("Hola mundo goroutine")
	done <- true
}

func calcSquares(number int, squareOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareOp <- sum
}

func calcCubes(number int, cubeOp chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}
	cubeOp <- sum
}

func sendData(sendCh chan<- int) {
	sendCh <- 33
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	log.Println("Channels in Golang")
	var a chan int // zero value of a channel is nil, a := make(chan int) is not nil
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T\n", a)
	}
	go func() {
		a <- 333 // write to a channel
	}()
	data := <-a // read from a channel
	fmt.Println("Data read from the channel is", data)
	done := make(chan bool)
	go hola(done)
	<-done

	number := 333
	sqrCh := make(chan int)
	cubeCh := make(chan int)
	go calcSquares(number, sqrCh)
	go calcCubes(number, cubeCh)
	squares, cubes := <-sqrCh, <-cubeCh
	fmt.Printf("Final output is %d and %d/n", squares, cubes)

	sendCh := make(chan<- int) // sendCh := make(chan int)
	go sendData(sendCh)
	// fmt.Println(<-sendCh) // this will throw error as sendCh is write only channel

	fmt.Println("main function done")
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Received", v, ok)
	}
	ch2 := make(chan int)
	go producer(ch2)
	for v := range ch2 {
		fmt.Println("Received", v)
	}
}
