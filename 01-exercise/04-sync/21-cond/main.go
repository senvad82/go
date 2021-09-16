package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)

	wg.Add(1)
	c.L.Lock()
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.

		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			c.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
	}()
	c.L.Unlock()
	// writes changes to sharedRsc
	c.L.Lock()
	sharedRsc["rsc1"] = "foo"
	c.Signal()
	c.L.Unlock()
	wg.Wait()
}
