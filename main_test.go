package main

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("Total:", multiConsumerProducer(10, 20))
}
