package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// mutex creates an exclusive shared memory
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
