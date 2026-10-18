package main

import "fmt"

// send-only channel
func producer(ch chan<- int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch) // only the sender closes
}

// receive-only channel
func consumer(ch <-chan int) {
	for v := range ch { // reads until the channel is closed
		fmt.Println("Got:", v)
	}
}

func main() {
	// unbuffered: send blocks until someone receives
	unbuffered := make(chan int)
	go producer(unbuffered)
	consumer(unbuffered)

	// buffered: send only blocks when the buffer is full
	buffered := make(chan string, 2)
	buffered <- "a"
	buffered <- "b"
	fmt.Println(len(buffered), cap(buffered)) // 2 2

	close(buffered)
	v, ok := <-buffered
	fmt.Println(v, ok) // a true
	<-buffered
	v, ok = <-buffered
	fmt.Printf("%q %v\n", v, ok) // "" false (closed and empty)
}
