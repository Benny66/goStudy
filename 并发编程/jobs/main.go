package main

import (
	"fmt"
	"sync"
)

type BlockingQueue struct {
	queue []interface{}
	mu    sync.Mutex
	cond  *sync.Cond
}

func NewBlockingQueue() *BlockingQueue {
	bq := &BlockingQueue{}
	bq.cond = sync.NewCond(&bq.mu)
	return bq
}

func (bq *BlockingQueue) Enqueue(item interface{}) {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	bq.queue = append(bq.queue, item)
	bq.cond.Signal()
}

func (bq *BlockingQueue) Dequeue() interface{} {
	bq.mu.Lock()
	defer bq.mu.Unlock()

	for len(bq.queue) == 0 {
		bq.cond.Wait()
	}

	item := bq.queue[0]
	bq.queue = bq.queue[1:]

	return item
}

func main() {
	bq := NewBlockingQueue()

	go func() {
		for i := 0; i < 10; i++ {
			bq.Enqueue(i)
		}
	}()

	go func() {
		for {
			item := bq.Dequeue()
			fmt.Println(item)
		}
	}()

	select {}
}
