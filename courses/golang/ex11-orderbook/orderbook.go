package orderbook

import "sort"

type Orderbook struct {
	RestingBids []*Order
	RestingAsks []*Order
}

func New() *Orderbook {
	return &Orderbook{
		RestingBids: make([]*Order, 0),
		RestingAsks: make([]*Order, 0),
	}
}

func (orderbook *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	trades := make([]*Trade, 0)
	var diff uint64

	switch order.Kind {
	case KindLimit:
		switch order.Side {
		case SideBid:
			for i, ask := range orderbook.RestingAsks {
				if order.Volume == 0 {
					break
				}

				if ask.Volume == 0 {
					orderbook.RestingAsks = append(orderbook.RestingAsks[:i], orderbook.RestingAsks[i+1:]...)
					continue
				}

				if ask.Price > order.Price {
					continue
				}

				if ask.Volume < order.Volume {
					diff = ask.Volume
				} else {
					diff = order.Volume
				}
				order.Volume -= diff
				ask.Volume -= diff

				trades = append(trades, &Trade{
					Bid:    order,
					Ask:    ask,
					Volume: diff,
					Price:  ask.Price,
				})
			}

			if order.Volume != 0 {
				orderbook.RestingBids = append(orderbook.RestingBids, order)
				sort.Slice(orderbook.RestingBids, func(i, j int) bool {
					return orderbook.RestingBids[i].Price > orderbook.RestingBids[j].Price
				})
			}
		case SideAsk:
			for i, bid := range orderbook.RestingBids {
				if order.Volume == 0 {
					break
				}

				if bid.Volume == 0 {
					orderbook.RestingBids = append(orderbook.RestingBids[:i], orderbook.RestingBids[i+1:]...)
					continue
				}

				if bid.Price < order.Price {
					continue
				}

				if bid.Volume < order.Volume {
					diff = bid.Volume
				} else {
					diff = order.Volume
				}
				order.Volume -= diff
				bid.Volume -= diff

				trades = append(trades, &Trade{
					Bid:    bid,
					Ask:    order,
					Volume: diff,
					Price:  bid.Price,
				})
			}

			if order.Volume != 0 {
				orderbook.RestingAsks = append(orderbook.RestingAsks, order)
				sort.Slice(orderbook.RestingAsks, func(i, j int) bool {
					return orderbook.RestingAsks[i].Price < orderbook.RestingAsks[j].Price
				})
			}
		}
	case KindMarket:
		switch order.Side {
		case SideBid:
			for i, ask := range orderbook.RestingAsks {
				if order.Volume == 0 {
					break
				}

				if ask.Volume == 0 {
					orderbook.RestingAsks = append(orderbook.RestingAsks[:i], orderbook.RestingAsks[i+1:]...)
					continue
				}

				if ask.Volume < order.Volume {
					diff = ask.Volume
				} else {
					diff = order.Volume
				}
				order.Volume -= diff
				ask.Volume -= diff

				trades = append(trades, &Trade{
					Bid:    order,
					Ask:    ask,
					Volume: diff,
					Price:  ask.Price,
				})
			}
		case SideAsk:
			for i, bid := range orderbook.RestingBids {
				if order.Volume == 0 {
					break
				}

				if bid.Volume == 0 {
					orderbook.RestingBids = append(orderbook.RestingBids[:i], orderbook.RestingBids[i+1:]...)
					continue
				}

				if bid.Volume < order.Volume {
					diff = bid.Volume
				} else {
					diff = order.Volume
				}
				order.Volume -= diff
				bid.Volume -= diff

				trades = append(trades, &Trade{
					Bid:    bid,
					Ask:    order,
					Volume: diff,
					Price:  bid.Price,
				})
			}
		}

		if order.Volume != 0 {
			return trades, order
		}
	}
	return trades, nil
}
