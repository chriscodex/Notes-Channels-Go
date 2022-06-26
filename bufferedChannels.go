package main

import (
	"fmt"
	"time"
)

func doSomething(i int) {
	fmt.Printf("Started %d", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Finished %d", i)

}

func main() {

}
