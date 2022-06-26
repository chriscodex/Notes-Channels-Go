package main

func generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i + 1
	}
	close(c)
}

func double(in <-chan int, out chan int) {
	for value := range in {
		out <- value
	}
	close(out)
}

func main() {

}
