package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	n  int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *Cache) Get(k string) string {
	c.mu.RLock() // many readers at the same time
	defer c.mu.RUnlock()
	return c.data[k]
}

func (c *Cache) Set(k, v string) {
	c.mu.Lock() // exclusive
	defer c.mu.Unlock()
	c.data[k] = v
}

func main() {
	var wg sync.WaitGroup

	counter := &Counter{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()
	fmt.Println("counter:", counter.n) // always 1000 (try without the mutex using -race)

	cache := &Cache{data: map[string]string{}}
	cache.Set("lang", "go")
	fmt.Println("cache:", cache.Get("lang"))
}
