package main

import (
	"fmt"
	"testing"
)

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrder1 := NewOrder(1.0, true)
	buyOrder2 := NewOrder(2.0, true)
	buyOrder3 := NewOrder(3.0, true)

	l.AddOrder(buyOrder1)
	l.AddOrder(buyOrder2)
	l.AddOrder(buyOrder3)

	if l.TotalVolume != 6.0 {
		t.Errorf("TotalVolume is not correct")
	}

	l.DeleteOrder(buyOrder2)

	if l.TotalVolume != 4.0 {
		t.Errorf("TotalVolume is not correct")
	}
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
