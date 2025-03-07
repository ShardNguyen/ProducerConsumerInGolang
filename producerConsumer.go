/*
Task: Write a program, where producers send values to consumers,
and calculate the sum of values the consumers has received
*/

/*
- Producer:
	ch: int channel
	val: int
	NewProducer(ch, value) : constructor
	SetValue(newValue)
	GetValue()
	SendValue()

- Consumer:
	ch: int channel
	NewConsumer(ch) : constructor
	ReceiveValue()

- Counter:
	numActiveProducer: int
	numActiveConsumer: int
*/

package main

// ----- Producer -----
type Producer struct {
	ch  chan int
	val int
}

func NewProducer(ch chan int, val int) *Producer {
	return &Producer{
		ch:  ch,
		val: val,
	}
}

func (p *Producer) SetValue(newVal int) {
	p.val = newVal
}

func (p *Producer) GetValue() int {
	return p.val
}

func (p *Producer) SendValue() {
	p.ch <- p.val
}

// ----- Consumer -----
type Consumer struct {
	ch chan int
}

func NewConsumer(ch chan int) *Consumer {
	return &Consumer{
		ch: ch,
	}
}

func (c *Consumer) ReceiveValue() {
	for {
		select {}
	}
}
