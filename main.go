package main

import (
	"fmt"
	"time"
)

func multiProducerConsumer(producerSize, consumerSize int) int {
	// Some channel variables
	messagesPerProducer := 5
	ch := make(chan int) // Channel for data transferring

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
		c.ReceiveValue()
	}

	time.Sleep(2 * time.Second)
	close(ch)
	return 0
}

func main() {
	multiProducerConsumer(5, 5)
	fmt.Println("Finished")
}
