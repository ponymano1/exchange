package main

import (
	"fmt"
	"sort"
	"time"
)

//limit (限价单) 核心是一个价格，一个order列表，一个总量
//order 里有总量size, 买卖方向bid, 时间戳timestamp，和一个指向limit的指针
//orderbook 核心是两个limit列表，一个是买单，一个是卖单
type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}

type Orders []*Order

func (orders Orders) Len() int {
	return len(orders)
}

func (orders Orders) Swap(i, j int) {
	orders[i], orders[j] = orders[j], orders[i]
}

func (orders Orders) Less(i, j int) bool {
	return orders[i].Timestamp < orders[j].Timestamp
}

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

func (order *Order) IsFilled() bool {
	return order.Size == 0.0
}

func (order *Order) String() string {
	return fmt.Sprintf("Order{Size: %f, Bid: %t, Timestamp: %d}", order.Size, order.Bid, order.Timestamp)
}

type Limits []*Limit

type ByBestAsk struct{ Limits }

func (a ByBestAsk) Len() int {
	return len(a.Limits)
}

func (a ByBestAsk) Swap(i, j int) {
	a.Limits[i], a.Limits[j] = a.Limits[j], a.Limits[i]
}

func (a ByBestAsk) Less(i, j int) bool {
	return a.Limits[i].Price < a.Limits[j].Price
}

type ByBestBid struct{ Limits }

func (a ByBestBid) Len() int {
	return len(a.Limits)
}

func (a ByBestBid) Swap(i, j int) {
	a.Limits[i], a.Limits[j] = a.Limits[j], a.Limits[i]
}

func (a ByBestBid) Less(i, j int) bool {
	return a.Limits[i].Price > a.Limits[j].Price
}

type Limit struct {
	Price       float64
	Orders      Orders
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
	sort.Sort(limit.Orders)
}

func (limit *Limit) MatchOrders(dst *Order) []Match {
	var (
		matchs      []Match
		ordersToDel []*Order
	)

	for _, order := range limit.Orders {
		if dst.IsFilled() {
			break
		}

		match := limit.matchOrder(dst, order)
		matchs = append(matchs, match)

		limit.TotalVolume -= match.SizeFilled

		if order.IsFilled() {
			ordersToDel = append(ordersToDel, order)
		}
	}

	for _, order := range ordersToDel {
		limit.DeleteOrder(order)
	}

	return matchs
}

/**
前提: o1, o2 是不同方向的order
*/
func (limit *Limit) matchOrder(o1, o2 *Order) Match {
	var (
		ask        *Order
		bid        *Order
		sizeFilled float64
	)
	if o1.Bid {
		bid = o1
		ask = o2
	} else {
		bid = o2
		ask = o1
	}

	if bid.Size > ask.Size {
		sizeFilled = ask.Size
		bid.Size -= ask.Size
		ask.Size = 0
	} else {
		sizeFilled = bid.Size
		ask.Size -= bid.Size
		bid.Size = 0
	}

	return Match{
		Ask:        ask,
		Bid:        bid,
		SizeFilled: sizeFilled,
		Price:      limit.Price,
	}
}

func (limit *Limit) String() string {
	return fmt.Sprintf("Limit{Price: %f, TotalVolume: %f, OrdersCntL %d}", limit.Price, limit.TotalVolume, len(limit.Orders))
}

type OrderBook struct {
	asks []*Limit
	bids []*Limit

	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewOrderBook() *OrderBook {
	return &OrderBook{
		asks:      []*Limit{},
		bids:      []*Limit{},
		AskLimits: map[float64]*Limit{},
		BidLimits: map[float64]*Limit{},
	}
}

/**
生成一个市价单
*/
func (ob *OrderBook) PlaceMarketOrder(order *Order) []Match {
	//logic
	return []Match{}
}

/**
生成一个限价单
*/
func (ob *OrderBook) PlaceLimitOrder(price float64, order *Order) {
	//logic
	ob.add(price, order)

}

func (ob *OrderBook) PlaceOrder(price float64, order *Order) []Match {
	//logic
	ob.add(price, order)
	return []Match{}
}

func (ob *OrderBook) add(price float64, order *Order) {
	if order.Bid {
		ob.addBid(price, order)
	} else {
		ob.addAsk(price, order)
	}
}

func (ob *OrderBook) addBid(price float64, order *Order) {
	limit, ok := ob.BidLimits[price]
	if !ok {
		limit = NewLimit(price)
		ob.BidLimits[price] = limit
		ob.bids = append(ob.bids, limit)
	}
	limit.AddOrder(order)
}

func (ob *OrderBook) addAsk(price float64, order *Order) {
	limit, ok := ob.AskLimits[price]
	if !ok {
		limit = NewLimit(price)
		ob.AskLimits[price] = limit
		ob.asks = append(ob.asks, limit)
	}
	limit.AddOrder(order)
}

func (ob *OrderBook) String() string {
	return fmt.Sprintf("OrderBook{asks: %v, bids: %v}", ob.asks, ob.bids)
}

func (ob *OrderBook) Asks() []*Limit {
	sort.Sort(ByBestAsk{ob.asks})
	return ob.asks
}

func (ob *OrderBook) Bids() []*Limit {
	sort.Sort(ByBestBid{ob.bids})
	return ob.bids
}
