package main

import "fmt"

func generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
	close(c)
}

func double(in <-chan int, out chan<- int) {
	for value := range in {
		out <- value * 2
	}
	close(out)
}

func print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {
	generators := make(chan int)
	doubles := make(chan int)
	go generator(generators)
	go double(generators, doubles)
	print(doubles)
}
