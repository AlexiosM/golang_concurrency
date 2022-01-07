package main

import "fmt"

// TODO: Build a Pipeline
// generator() -> square() -> print

// generator - convertes a list of integers to a channel
func generator(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for r := range in {
			out <- r * r
		}
		close(out)
	}()
	return out
}

func main() {
	// set up the pipeline

	for i := range square(generator(2, 543, 76, 4568, 76, 235456)) {
		fmt.Printf("%v ", i)
	}

	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.

	// ch1 := generator(2, 3, 4, 5, 8, 2343, 34, 2124, 45756, 3, 346, 7467, 5, 8, 6, 9111, 72, 97)
	// ch2 := square(ch1)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	for i := range ch2 {
	// 		fmt.Printf("%v  ", i)
	// 	}
	// }()
	// wg.Wait()
}
