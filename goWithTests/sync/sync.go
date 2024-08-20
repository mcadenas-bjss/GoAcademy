package sync

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// I personally don't like this practice.
// Go is not object orientated so use a regular declaration and dereference where necessary
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
