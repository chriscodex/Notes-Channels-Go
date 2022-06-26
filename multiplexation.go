package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	go sleeping(2*time.Second, c1, 2)
	go sleeping(4*time.Second, c2, 4)
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}

func sleeping(i time.Duration, c chan int, param int) {
	time.Sleep(i)
	c <- param
}
