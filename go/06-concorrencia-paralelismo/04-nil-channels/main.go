package main

import "fmt"

func send(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}

// merge two channels: a closed channel becomes nil so its case is never selected again
func main() {
	ch1 := send(1, 2, 3)
	ch2 := send(10, 20)

	for ch1 != nil || ch2 != nil {
		select {
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			fmt.Println("ch1:", v)
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			fmt.Println("ch2:", v)
		}
	}
}
