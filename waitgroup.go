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

func doSomethingWaitGroup(i int) {
	fmt.Printf("Started %d", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}

func main() {
	for i := 0; i < 3; i++ {
		go doSomethingWaitGroup(i)
	}

}
