package main

import "fmt"

func generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
	close(c)
}

func double(in <-chan int, out chan int) {
	for value := range in {
		out <- value
	}
	close(out)
}

func print(c <-chan int) {
	for value := range c {
		fmt.Println(value)
	}
}

func main() {

}
