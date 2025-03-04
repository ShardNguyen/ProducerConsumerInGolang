package main

import (
	"fmt"
	"sync"
	"time"
)

func multiConsumerProducer(producerSize, consumerSize int) {
	ch := make(chan string) // A channel for communication between producer and consumer
	var wg sync.WaitGroup   // Wait Group

	// Start multiple producers
	for i := 0; i < producerSize; i++ {
		wg.Add(1) // Add one go function to the wait group
		go producer(i, ch, &wg)
	}

	// Start multiple consumers
	for i := 0; i < consumerSize; i++ {
		wg.Add(1) // Add one go function to the wait group
		go consumer(i, ch, &wg)
	}

	wg.Wait() // Wait until every go functions finish its job
	close(ch)
}

func producer(index int, ch chan string, wg *sync.WaitGroup) {
	// Each producer will send out 5 messages (As an example)
	for i := 0; i < 10; i++ {
		// Sending signal away and block the current iteration (But the next iteration can start somehow)
		ch <- fmt.Sprintf("Producer %v send %v", index, i)
	}
	wg.Done()
}

func consumer(index int, ch chan string, wg *sync.WaitGroup) {
	done := false
	for !done {
		select {
		// If the consumer can catch the channel
		case msg, ok := <-ch:
			if !ok {
				done = true
			}
			fmt.Printf("Consumer %v Received: %s\n", index, msg)
		// Timeout in case the producer doesn't produce any more responses (In this example, 1000ms)
		case <-time.After(1000 * time.Millisecond):
			wg.Done()
		}
	}
	wg.Done()
}

func main() {
	multiConsumerProducer(20, 2)
}
