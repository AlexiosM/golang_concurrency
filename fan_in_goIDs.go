package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

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

func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idFields := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))
	fmt.Println(idFields)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("goroutine id: %d, read value %d\n§", goid(), n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func main() {
	in := generator(2, 3, 4, 5, 6, 574, 235, 6)

	ch1 := square(in)
	ch2 := square(in)
	ch3 := square(in)
	ch4 := square(in)

	time.Sleep(2 * time.Second)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)
	fmt.Println(<-ch3)
	fmt.Println(<-ch4)
}
