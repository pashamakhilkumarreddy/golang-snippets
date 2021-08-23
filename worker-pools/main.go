package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digits(number int) int {
	sum, num := 0, number
	for num != 0 {
		digit := num % 10
		sum += digit
		num = num / 10
	}
	time.Sleep(1 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomno)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randonNo := rand.Intn(999)
		job := Job{i, randonNo}
		jobs <- job
	}
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func main() {
	log.Println("Wait Groups and Worker Pools in Golang")
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
	startTime := time.Now()
	noOfJobs := 100
	fmt.Println("Worker Pool started executing at", startTime)
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done

	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("Worker Pool finished at", endTime, " in ", diff)
}
