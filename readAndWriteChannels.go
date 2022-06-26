package main

func generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
}

func main() {

}
