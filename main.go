// Task: Replace waitgroup with channel, is that ok?
/*

 */

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
)

func multiConsumerProducer(producerSize, consumerSize int) int {
	ch := make(chan int)           // A channel for communication between producer and consumer
	stopCh := make(chan struct{})  // An additional channel to stop data
	toStopCn := make(chan bool)    // Another channel to notify the moderator to close the stopCh channel
	var wgConsumers sync.WaitGroup // Wait Group to indicate how many concurrent RECEIVING functions that needs waiting to be done
	sum := 0
	var mu sync.RWMutex

	// Start one moderator
	go moderator(stopCh, toStopCn)

	// Start multiple producers
	for i := 0; i < producerSize; i++ {
		go producer(i, producerSize, ch, stopCh, toStopCn)
	}

	// Start multiple consumers
	wgConsumers.Add(consumerSize)
	for i := 0; i < consumerSize; i++ {
		go consumer(i, ch, stopCh, toStopCn, &wgConsumers, &mu, &sum)
	}

	wgConsumers.Wait() // Wait until every consumers functions finish their job
	close(ch)
	return sum
}

func moderator(stopCh chan struct{}, toStopCn chan bool) {
	<-toStopCn
	close(stopCh)
}

func producer(index int, prodSize int, ch chan int, stopCh chan struct{}, toStopCn chan bool) {
	// Each producer will send out 5 messages (As an example)
	for i := 0; i < 5; i++ {
		// Logical condition when producer doesn't want to produce any more data
		if index+1 == prodSize && i == 4 {
			select {
			// Send a signal to the moderator for ending the data tranmission
			case toStopCn <- true:
				fmt.Println("Stopped by producer", index, "at msg", i)
			default:
			}
		}

		randNum := rand.IntN(100)

		select {
		// Stop sending messages when receiving the stop signal
		case <-stopCh:
			return
		// Sending signal away and block the thread of current iteration
		case ch <- randNum:
			fmt.Println("Produced", randNum)
		}
	}
}

func consumer(index int, ch chan int, stopCh chan struct{}, toStopCn chan bool, wg *sync.WaitGroup, mu *sync.RWMutex, sum *int) {
	defer wg.Done()
consumerLoop:
	for {
		select {
		// Case when receiever caught the message
		case msg := <-ch:
			// Mutex is needed to synchronize the sum that needs reading
			mu.Lock()
			fmt.Println("Received", msg)
			*sum += msg
			mu.Unlock()

			// Logical condition when consumer doesn't want to consume any more data
			if index == 7 {
				select {
				// Send a signal to the moderator for ending the data tranmission
				case toStopCn <- true:
					fmt.Println("Consumer", index, "asked to stop")
				default:
				}
				break consumerLoop
			}
		case <-stopCh:
			break consumerLoop
		}
	}
}

func main() {
	fmt.Println("Total:", multiConsumerProducer(10, 20))
}
