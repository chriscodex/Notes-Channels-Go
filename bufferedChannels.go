package main

import (
	"fmt"
	"sync"
	"time"
)

func doSomething(i int, c chan int, wg *sync.WaitGroup) {
	fmt.Printf("Started %d\n", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished %d\n", i)
	<-c
	defer wg.Done()
}

func main() {
	c := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		c <- 1
		wg.Add(1)
		go doSomething(i, c, &wg)
	}
	<-c
	wg.Wait()
}
