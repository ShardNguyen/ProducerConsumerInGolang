// Task: Change bool condition on loop to channel-based condition kind of loop
/*
	- Current solution: Make a moderator role

	- Moderator will be the only channel to trigger the toStop channel
	- toStop channel will notify the moderator to close Stop channel, meaning the Stop channel will be triggered in the "select" case
		+ Sending data to a closed channel will cause a panic
		+ But receiving data from a closed channel will be ok (Meaning <- closedChannel will still trigger the "select")
	- The Stop channel is used for breaking the loop in both consumer and producer
*/

package main

import (
	"fmt"
	"sync"
)

func multiConsumerProducer(producerSize, consumerSize int) {
	ch := make(chan string)        // A channel for communication between producer and consumer
	stopCh := make(chan struct{})  // An additional channel to stop data
	toStopCn := make(chan bool)    // Another channel to notify the moderator to close the stopCh channel
	var wgConsumers sync.WaitGroup // Wait Group to indicate how many concurrent RECEIVING functions that needs waiting to be done

	// Start one moderator
	go moderator(stopCh, toStopCn)

	// Start multiple producers
	for i := 0; i < producerSize; i++ {
		go producer(i, producerSize, ch, stopCh, toStopCn)
	}

	// Start multiple consumers
	for i := 0; i < consumerSize; i++ {
		wgConsumers.Add(1) // Indicate that Wait Group needs to wait for one more function
		go consumer(i, ch, stopCh, toStopCn, &wgConsumers)
	}

	wgConsumers.Wait() // Wait until every consumers functions finish their job
	close(ch)
}

func moderator(stopCh chan struct{}, toStopCn chan bool) {
	<-toStopCn
	close(stopCh)
}

func producer(index int, prodSize int, ch chan string, stopCh chan struct{}, toStopCn chan bool) {
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

		select {
		// Stop sending messages when receiving the stop signal
		case <-stopCh:
			return
		// Sending signal away and block the thread of current iteration
		case ch <- fmt.Sprintf("Producer %v send %v", index, i):
		}
	}
}

func consumer(index int, ch chan string, stopCh chan struct{}, toStopCn chan bool, wg *sync.WaitGroup) {
consumerLoop:
	for {
		select {
		// Case when receiever caught the message
		case msg := <-ch:
			fmt.Printf("Consumer %v Received: %s\n", index, msg)
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
	wg.Done()
}

func main() {
	multiConsumerProducer(20, 10)
}
