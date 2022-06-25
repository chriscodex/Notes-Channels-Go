package main

import (
	"fmt"
)

func goFunc(c chan int) {
	c <- 1
}
func goFunc2(c chan int) {
	fmt.Println(c)
	c <- 2
	c <- 4
}

func main() {
	c := make(chan int)
	go goFunc(c)
	fmt.Println(<-c)
}
