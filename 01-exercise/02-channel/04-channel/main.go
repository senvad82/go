package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(c1 chan<- int) {

	c1 <- 10
	// send message on ch1
}

func relayMsg(c1 <-chan int, c2 chan<- int) {

	m := <-c1
	c2 <- m
	//close(c1)
	// recv message on ch1
	// send it on ch2
}

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)
	// create ch1 and ch2

	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	//fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	// spine goroutine genMsg and relayMsg

	// recv message on ch2
}
