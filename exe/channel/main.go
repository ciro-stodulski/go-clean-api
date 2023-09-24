package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go counter(naturals)
	go squarer(squares, naturals)

	printer(squares)

}

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	fmt.Printf("counter value: %v\n", <-in)
}

// gopl.io/ch8/cake
// simular confeitaria, concorrencia
