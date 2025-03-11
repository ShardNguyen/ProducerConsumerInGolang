/*
- Consumer:
	ch: int channel
	val: int

	NewConsumer(ch) : constructor
	ReceiveValue()
*/

package main

type Consumer struct {
	ch  chan int
	val int
}

func NewConsumer(ch chan int) *Consumer {
	return &Consumer{
		ch: ch,
	}
}

func (c *Consumer) ReceiveValue(counter *Counter) {
	for {
		select {
		case c.val = <-c.ch:
			// fmt.Println("Received", c.val)
			counter.AddConsumerMessage(c.val)
		}
	}
}
