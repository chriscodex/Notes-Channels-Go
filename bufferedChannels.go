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
	defer wg.Done()
}

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go doSomething(i, c, &wg)
		wg.Add(1)
	}
	wg.Wait()
}
