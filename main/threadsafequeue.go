package main

import (
	"fmt"
	"sync"
)

type ConcurrentQueue struct {
	queue []interface{}
	mu    sync.Mutex
}

var wgE sync.WaitGroup

func (q *ConcurrentQueue) enqueue(object interface{}) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.queue = append(q.queue, object)
}

func (q *ConcurrentQueue) dequeue() interface{} {
	if len(q.queue) == 0 {
		return 0
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func main() {
	q1 := ConcurrentQueue{
		queue: make([]interface{}, 0),
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
