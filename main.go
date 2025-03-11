// To do: Fix race condition

package main

import (
	"fmt"
	"time"
)

func multiProducerConsumer(producerSize, consumerSize int) int {
	messagesPerProducer := 5
	ch := make(chan int) // Channel for data transferring
	var counter Counter

	// Create and start multiple producers
	for i := 0; i < producerSize; i++ {
		go func(i int) {
			p := NewProducer(ch)

			for j := 0; j < messagesPerProducer; j++ {
				go p.SendValue(i*messagesPerProducer + j)
			}
		}(i)
	}

	// Create and start multiple consumers
	for i := 0; i < consumerSize; i++ {
		c := NewConsumer(ch)
		go c.ReceiveValue(&counter)
	}

	time.Sleep(3 * time.Second)
	close(ch)
	return counter.ReturnSum()
}

func main() {
	sum := multiProducerConsumer(25, 1)
	fmt.Println("Finished with sum being:", sum)
}
