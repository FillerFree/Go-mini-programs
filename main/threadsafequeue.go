package main

import (
	"fmt"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
	mu    sync.Mutex
}

var wgE sync.WaitGroup

func (q *ConcurrentQueue) enqueue(num int32) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, num)
}

func (q *ConcurrentQueue) dequeue() int32 {
	if len(q.queue) == 0 {
		return 0
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func main() {
	q1 := ConcurrentQueue{
		queue: make([]int32, 0),
	}
	for i := 0; i < 100000; i += 1 {
		wgE.Add(1)
		go func() {
			q1.enqueue(1)
			wgE.Done()
		}()
	}
	wgE.Wait()
	fmt.Println(len(q1.queue))
}
