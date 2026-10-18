package main

import (
	"context"
	"fmt"
	"time"
)

func slowWork(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Work finished")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	// WithTimeout: cancels automatically after 1s
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	fmt.Println("timeout:", slowWork(ctx)) // context deadline exceeded

	// WithCancel: cancels manually
	ctx2, cancel2 := context.WithCancel(context.Background())
	go func() {
		time.Sleep(500 * time.Millisecond)
		cancel2()
	}()

	fmt.Println("cancel:", slowWork(ctx2)) // context canceled

	// WithValue: request scoped data
	type key string
	ctx3 := context.WithValue(context.Background(), key("user"), "miguel")
	fmt.Println("value:", ctx3.Value(key("user")))
}
