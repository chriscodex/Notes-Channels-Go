package main

import "time"

func main() {

}

func sleeping(i time.Duration, c chan int) {
	time.Sleep(i)
}
