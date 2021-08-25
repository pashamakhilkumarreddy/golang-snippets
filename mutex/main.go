package main

import (
	"fmt"
	"log"
	"sync"
)

var (
	x = 0
	y = 1000
)

func increment(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func decrementCh(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	y = y - 1
	<-ch
	wg.Done()
}

func main() {
	log.Println("Mutex in Golang")

	var wg sync.WaitGroup
	var mu sync.Mutex

	var wgCh sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mu)
	}
	wg.Wait()
	fmt.Printf("Final value of x is %d\n", x)

	for i := 0; i < 1000; i++ {
		wgCh.Add(1)
		go decrementCh(&wgCh, ch)
	}
	wgCh.Wait()

	fmt.Printf("Final value of y is %d\n", y)
}
