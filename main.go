// Task: Log the timeout case and check if the time out is reached
/*
	- All of the consumers will reach timeout
	- I believe this has to be related to the "select" part on the consumers' side
	- Because although when producer no longer sends the message, the consumer will continue on waiting, making case msg := <-ch unable to happen
	- And that will cause the timeout case to happen
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func multiConsumerProducer(producerSize, consumerSize int) {
	ch := make(chan string) // A channel for communication between producer and consumer
	var wg sync.WaitGroup   // Wait Group to indicate how many concurrent functions that needs waiting to be done

	// Start multiple producers
	for i := 0; i < producerSize; i++ {
		wg.Add(1) // Indicate that Wait Group needs to wait for one more function
		go producer(i, ch, &wg)
	}

	// Start multiple consumers
	for i := 0; i < consumerSize; i++ {
		wg.Add(1) // Indicate that Wait Group needs to wait for one more function
		go consumer(i, ch, &wg)
	}

	wg.Wait() // Wait until every go functions finish their job
	close(ch)
}

func producer(index int, ch chan string, wg *sync.WaitGroup) {
	// Each producer will send out 5 messages (As an example)
	for i := 0; i < 5; i++ {
		// Sending signal away and block the current iteration (But the next iteration can start somehow?)
		ch <- fmt.Sprintf("Producer %v send %v", index, i)
	}
	// Once all of the producer's messages are received, the job is done for the producer
	wg.Done()
}

func consumer(index int, ch chan string, wg *sync.WaitGroup) {
	done := false
	for !done {
		// The loop is blocked until one of the two scenarios happens
		select {
		// Case when receiever caught the message
		case msg := <-ch:
			fmt.Printf("Consumer %v Received: %s\n", index, msg)
		// Timeout in case the producer doesn't produce any more responses (In this example, 1000ms)
		case <-time.After(1000 * time.Millisecond):
			fmt.Printf("Consumer %v Timeout!\n", index)
			done = true
		}
	}
	wg.Done()
}

func main() {
	multiConsumerProducer(2, 10)
}
