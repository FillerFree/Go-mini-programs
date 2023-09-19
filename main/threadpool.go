package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ConcurrentPrint struct {
	name     int32
	function func()
}

func CreatePool(size int) {
	q := ConcurrentQueue{
		queue: make([]interface{}, 0),
	}

	for i := 0; i < size; i += 1 {
		name := int32(rand.Int())
		q.enqueue(ConcurrentPrint{
			name: name,
			function: func() {
				go func() {
					fmt.Println(name)
					time.Sleep(time.Second)
				}()
			},
		})
	}

	for i := 0; i < size; i += 1 {
		go func() {
			val := ConcurrentPrint(q.dequeue())
		}()
	}
}
