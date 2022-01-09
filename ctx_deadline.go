package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func main() {

	// TODO: set deadline for goroutine to return computational result.

	deadline := time.Now().Add(10 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel() // if the work is completed before the deadline, free the resources

	compute := func() <-chan data {
		ch := make(chan data)
		go func() {
			defer close(ch)

			deadline, ok := ctx.Deadline() // get the time when the work should be canceled
			if ok {                        // if deadline has been set
				if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
					fmt.Println("not sufficient time given...terminating")
					return
				}
			}

			// Simulate work.
			time.Sleep(50 * time.Millisecond)

			// Report result.
			select {
			case ch <- data{"123"}:
			case <-ctx.Done():
				return
			}
		}()
		return ch
	}

	// Wait for the work to finish. If it takes too long move on.
	ch := compute()
	d, ok := <-ch
	if ok { // if it is due to send operation
		fmt.Printf("work complete: %s\n", d)
	}
}
