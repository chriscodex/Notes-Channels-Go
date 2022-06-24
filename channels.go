package main

import (
	"fmt"
)

func goFunc(c chan int) {
	c <- 3
	c <- 4
}
func goFunc2(c chan int) {
	c <- 5
	c <- 6
}

func main() {
	c := make(chan int, 4)
	go goFunc(c)
	go goFunc2(c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
