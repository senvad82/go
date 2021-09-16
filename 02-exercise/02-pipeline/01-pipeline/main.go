package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	ch := make(chan int)
	go func() {
		for _, value := range nums {
			ch <- value
		}
		close(ch)
	}()

	return ch

}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(ch <-chan int) <-chan int {
	ch2 := make(chan int)

	go func() {
		for value := range ch {
			ch2 <- value * value
		}
		close(ch2)
	}()

	return ch2
}

func main() {
	// set up the pipeline

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	for res := range square(square(generator(3, 4))) {
		fmt.Println(res)
	}

}
