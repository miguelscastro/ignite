package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int) {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker", id, "shutting down")
			time.Sleep(300 * time.Millisecond) // finishing in-flight work
			return
		case <-ticker.C:
			fmt.Println("worker", id, "working")
		}
	}
}

// run and press Ctrl+C
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ctx, i)
		}()
	}

	<-ctx.Done()
	fmt.Println("signal received, waiting for workers...")
	wg.Wait()
	fmt.Println("shutdown complete")
}
