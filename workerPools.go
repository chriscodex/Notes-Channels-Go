package main

import "fmt"

func main() {
	tasks := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	numberWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))
	// Go routines
	for i := 0; i < numberWorkers; i++ {
		go worker(i+1, jobs, results)
	}
	// Sending numbers
	for _, value := range tasks {
		jobs <- value
	}
	close(jobs)
	// Show values
	for r := 0; r < len(tasks); r++ {
		<-results
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker id: %d, started double of %d\n", id, job)
		dou := doubles(job)
		fmt.Printf("Worker id: %d, job %d and double %d\n", id, job, dou)
		results <- dou
	}
}

func doubles(n int) int {
	return n * 2
}
