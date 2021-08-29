package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("./assets/concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err := fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	log.Println("Writing to Files in Golang")
	func() {
		f, err := os.Create("./assets/names.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		l, err := f.WriteString("Gwen Stacy ")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		fmt.Println(l, "write to the file in bytes is successfully")
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	func() {
		f2, err := os.Create("./assets/mensage.txt")
		if err != nil {
			fmt.Println(err)
			return
		}
		b2 := []byte{104, 111, 108, 97, 32, 119, 111, 114, 108, 100}
		d2, err := f2.Write(b2)
		if err != nil {
			fmt.Println(err)
			f2.Close()
			return
		}
		fmt.Println(d2, "write to the file in bytes in successfully")
		err = f2.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()

	func() {
		f, err := os.Create("./assets/linea.txt")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		d3 := []string{"Hola Mundo!", "Como estas?", "Bienvenido a casa"}
		for _, v := range d3 {
			fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("write to file is successfull")
	}()

	func() {
		f, err := os.Create("./assets/amor.txt")
		if err != nil {
			fmt.Println(err)
			f.Close()
			return
		}
		d4 := []string{"Hola Gwen Stacy", "Te quiero"}
		for _, v := range d4 {
			fmt.Fprintln(f, v)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("write to file is successful")

		func() {
			f2, err := os.OpenFile("./assets/amor.txt", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			line := "Te quiero mucho"
			_, err = fmt.Fprintln(f2, line)
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f2.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("append to file is successfully")
		}()
	}()

	data, done := make(chan int), make(chan bool)
	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	if <-done {
		fmt.Println("write to file concurrently is successfully")
	} else {
		fmt.Println("write to file concurrently failed")
	}
}
