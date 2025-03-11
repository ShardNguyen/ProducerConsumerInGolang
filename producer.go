/*
Task: Write a program, where producers send values to consumers,
and calculate the sum of values the consumers has received
*/

/*
- Producer:
	ch: int channel
	NewProducer(ch) : constructor
	SendValue(val)
*/

package main

import "fmt"

// ----- Producer -----
type Producer struct {
	ch chan int
}

func NewProducer(ch chan int) *Producer {
	return &Producer{
		ch: ch,
	}
}

func (p *Producer) SendValue(val int) {
	p.ch <- val
	fmt.Println("Sent value", val)
}
