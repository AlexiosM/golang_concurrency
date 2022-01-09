package main

import (
	"fmt"
)

// ByteCounter type
type ByteCounter int

// TODO: Implement Write method for ByteCounter
// to count the number of bytes written.
func (bs *ByteCounter) Write(p []byte) (int, error) {
	*bs += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var b ByteCounter
	fmt.Fprintf(&b, "hello world")
	fmt.Println(b)
}
