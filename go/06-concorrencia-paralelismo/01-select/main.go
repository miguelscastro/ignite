package main

import (
	"fmt"
	"time"
)

func read(ch <-chan int) {
	for {
		select {
		case value := <-ch:
			fmt.Println("Got something: ", value)
			return
		default:
			fmt.Println("Nothing here yet")
		}
	}
}

// timeout: select blocks until a case is ready, time.After gives up after 1s
func timeoutExample() {
	ch := make(chan int)

	select {
	case v := <-ch:
		fmt.Println("Got:", v)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}

// select inside a for, stopped by a done channel and by a closed channel
func loopExample() {
	data := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < 3; i++ {
			data <- i
		}
		close(done)
	}()

	for {
		select {
		case v := <-data:
			fmt.Println("Data:", v)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

func main() {
	timeoutExample()
	loopExample()

	ch := make(chan int)

	go read(ch)

	time.Sleep(5 * time.Second)
	ch <- 10
}
