package main

import "fmt"

func main() {

}

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker id: %d, started double of %d\n", id, job)
		dou := doubles(job)
		fmt.Printf("Worker id: %d, job %d and double %d", id, job, dou)
		results <- dou
	}
}

func doubles(n int) int {
	return n * 2
}
