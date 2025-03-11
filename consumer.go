/*
- Consumer:
	ch: int channel
	val: int

	NewConsumer(ch) : constructor
	ReceiveValue()
*/

package main

import "fmt"

type Consumer struct {
	ch  chan int
	val int
}

func NewConsumer(ch chan int) *Consumer {
	return &Consumer{
		ch: ch,
	}
}

func (c *Consumer) ReceiveValue() {
	for {
		select {
		case c.val = <-c.ch:
			fmt.Println("Received", c.val)
		}
	}
}
