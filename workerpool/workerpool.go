package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	start := time.Now()

	const numJobs = 10
	const numWorkers = 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	resultSlice := []int{}
	for a := 1; a <= numJobs; a++ {
		resultSlice = append(resultSlice, <-results)
	}
	fmt.Println(len(resultSlice), " jobs finished in ", time.Now().Sub(start).Seconds())
}