package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(3) // at most 3 goroutines at the same time

	for i := 1; i <= 5; i++ {
		g.Go(func() error {
			select {
			case <-time.After(time.Duration(i) * 200 * time.Millisecond):
				if i == 3 {
					return errors.New("task 3 failed")
				}
				fmt.Println("task", i, "ok")
				return nil
			case <-ctx.Done(): // cancelled because another task failed
				fmt.Println("task", i, "cancelled")
				return ctx.Err()
			}
		})
	}

	if err := g.Wait(); err != nil { // first non-nil error
		fmt.Println("error:", err)
	}
}
