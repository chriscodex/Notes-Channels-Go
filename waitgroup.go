package main

import (
	"fmt"
	"time"
)

func doSomethingChannels(c chan struct{}) {
	fmt.Println("Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
	c <- struct{}{}
}

func doSomethingWaitGroup() {
	fmt.Println("Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}

func main() {
	c := make(chan struct{})
	for i := 0; i < 3; i++ {
		go doSomething(c)
	}
	for i := 0; i < 3; i++ {
		fmt.Println(<-c)
	}

}
