package main

import "sync"

type Counter struct {
	numMessagesProduced int
	numMessagesConsumed int
	sum                 int
	mu                  sync.RWMutex
}

func (count *Counter) AddProducerMessage() {
	count.mu.Lock()
	count.numMessagesProduced++
	count.mu.Unlock()
}

func (count *Counter) AddConsumerMessage(val int) {
	count.mu.Lock()
	count.numMessagesConsumed++
	count.sum += val
	count.mu.Unlock()
}

func (count *Counter) ReturnSum() int {
	return count.sum
}
