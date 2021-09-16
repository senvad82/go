package main

import "fmt"

func main() {

	ch := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			ch <- i
			// TODO: send iterator over channel
		}
		close(ch)
	}()

	for c := range ch {
		fmt.Println(c)
	}
	// TODO: range over channel to recv values

}
