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
	defer wg.Done()
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished %d\n", i)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doSomethingWaitGroup(i, &wg)
	}
	wg.Wait()
}
