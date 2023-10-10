package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	order := NewOrder(1.0, true)
	l.AddOrder(order)

	fmt.Println(l)

}

func TestOrderbook(t *testing.T) {
	ob := &OrderBook{}

	// Test that the OrderBook struct is initialized correctly
	if ob == nil {
		t.Errorf("OrderBook struct is nil")
	}

	// Add more tests here as needed
}
