package main

import "time"

func doSomething() {
	time.Sleep(2 * time.Second)
}

func main() {
	for i := 0; i < 10; i++ {
		go doSomething()
	}
}
