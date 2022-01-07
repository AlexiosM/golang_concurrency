package main

import "fmt"

// TODO: Implement relaying of message with Channel Direction

func genMsg(ch chan<- string) {
	// send message on ch1
	ch <- "Hello I'm Alex"
}

func relayMsg(in <-chan string, out chan<- string) {
	// recv message on ch1
	m := <-in
	// send it on ch2
	out <- m
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)
	// spine goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)
	// recv message on ch2
	fmt.Println(<-ch2)
}
