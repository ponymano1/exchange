package main

import (
	"fmt"
	"reflect"
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
	orderBook := NewOrderBook()
	buyOrder1 := NewOrder(100, true)
	buyOrder2 := NewOrder(200, true)
	buyOrder3 := NewOrder(100, true)
	orderBook.PlaceOrder(100, buyOrder1)
	orderBook.PlaceOrder(200, buyOrder2)
	orderBook.PlaceOrder(100, buyOrder3)

	for i := 0; i < len(orderBook.bids); i++ {
		fmt.Println(orderBook.bids[i].String())
	}

	if orderBook.bids[0].TotalVolume != 200 {
		t.Errorf("TotalVolume is not correct")
	}

	if orderBook.bids[1].TotalVolume != 200 {
		t.Errorf("TotalVolume is not correct")
	}

	if len(orderBook.bids[0].Orders) != 2 {
		t.Errorf("Order count is not correct")
	}
}

func TestOrderBook_PlaceOrder(t *testing.T) {
	type args struct {
		price float64
		order *Order
	}
	tests := []struct {
		name string
		ob   *OrderBook
		args args
		want []Match
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ob.PlaceOrder(tt.args.price, tt.args.order); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OrderBook.PlaceOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("Expected %v, got %v", a, b)
	}

}

func TestOrderBook_PlaceLimitOrder(t *testing.T) {
	ob := NewOrderBook()

	sellOrder := NewOrder(10, false)
	ob.PlaceLimitOrder(100_000, sellOrder)

	assert(t, len(ob.asks), 1)
}
