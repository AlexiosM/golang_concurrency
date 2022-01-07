package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup
	mu := sync.Mutex{}
	cond := sync.NewCond(&mu)

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc1"])
		cond.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		//TODO: suspend goroutine until sharedRsc is populated.
		cond.L.Lock()
		for len(sharedRsc) == 0 {
			//time.Sleep(1 * time.Millisecond)
			cond.Wait()
		}

		fmt.Println(sharedRsc["rsc2"])
		cond.L.Unlock()
	}()

	// writes changes to sharedRsc
	cond.L.Lock()
	sharedRsc["rsc1"] = "foo"
	sharedRsc["rsc2"] = "bar"
	cond.Broadcast()
	cond.L.Unlock()

	wg.Wait()
}
