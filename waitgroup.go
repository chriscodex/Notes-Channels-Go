package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomethingChannels(c chan struct{}) {
	fmt.Println("Started")
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
	c <- struct{}{}
}

func doSomethingWaitGroup(i int, wg *sync.WaitGroup) {
	fmt.Printf("Started %d", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doSomethingWaitGroup(i, &wg)
	}

}
