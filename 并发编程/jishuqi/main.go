package main

import (
	"fmt"
	"sync"
)

// 实现一个并发安全的计数器，可以支持多个goroutine同时读写。
type Counter struct {
	value int
	mutex sync.Mutex
}

func (c *Counter) Increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value++
}

func (c *Counter) Decrement() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.value--
}

func (c *Counter) GetValue() int {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.value
}

func main() {
	counter := Counter{}

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() { wg.Done() }()
			counter.Increment()
		}()
	}

	wg.Wait()

	fmt.Println(counter.GetValue())
}
