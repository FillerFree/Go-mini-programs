package main

type Bucket struct {
	name     string
	lastUsed int
	capacity int
}

var rateLimit = 3

func (b *Bucket) refill(currentMinute int) {
	b.capacity += currentMinute - b.lastUsed*3
	b.capacity = b.capacity % (rateLimit + 1)
}

func (b *Bucket) deduct() {
	if b.capacity == 0 {
		return
	}
	b.capacity -= 1
}

func (b *Bucket) isRateLimited(currentMinute int) bool {
	b.refill(currentMinute)
	b.lastUsed = currentMinute
	b.deduct()
	if b.capacity > 0 {
		return false
	}
	return true
}
