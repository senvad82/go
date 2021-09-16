package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	for {
		select {
		case g1 := <-ch1:
			fmt.Println(g1)
		case g2 := <-ch2:
			fmt.Println(g2)
		case <-time.After(4 * time.Second):
			fmt.Println("timeout")
		}
	}

	// TODO: multiplex recv on channel - ch1, ch2

}
