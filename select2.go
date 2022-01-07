package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "one"
	}()

	// TODO: implement timeout for recv on channel ch
	select {
	case r := <-ch:
		fmt.Println("Received: ", r)
	case tout := <-time.After(3 * time.Second):
		fmt.Printf("Timed out after %v seconds.\n", tout)
	}

	close(ch)
}
