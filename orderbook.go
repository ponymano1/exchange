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
	order.Limit = limit
	limit.Orders = append(limit.Orders, order)
	limit.TotalVolume += order.Size
}

func (limit *Limit) DeleteOrder(o *Order) {
	len := len(limit.Orders)
	for i := 0; i < len; i++ {
		if limit.Orders[i] == o {
			limit.Orders[i], limit.Orders[len-1] = limit.Orders[len-1], limit.Orders[i]
			limit.Orders = limit.Orders[0 : len-1]
			break
		}
	}

	o.Limit = nil
	limit.TotalVolume -= o.Size
}

type OrderBook struct {
	Ask []*Limit
	Bid []*Limit
}
