package main

import (
	"fmt"
	"time"
)

type Order struct {
	Size      float64
	Bid       bool
	Limit     *Limit
	Timestamp int64
}

func NewOrder(size float64, bid bool) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

func (order *Order) String() string {
	return fmt.Sprintf("Order{Size: %f, Bid: %t, Timestamp: %d}", order.Size, order.Bid, order.Timestamp)
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

func (limit *Limit) AddOrder(order *Order) {
	limit.Orders = append(limit.Orders, order)
}

type OrderBook struct {
	Ask []*Limit
	Bid []*Limit
}
