package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	fmt.Println("worker", id, "done")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // before starting the goroutine
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()
	fmt.Println("all workers finished")
}
