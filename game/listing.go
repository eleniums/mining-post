package game

type Listing struct {
	Resource

	Quantity  int64   `json:"quantity"`
	BuyPrice  float64 `json:"buy_price"`
	SellPrice float64 `json:"sell_price"`

	quantityRangeLow  int64   // lowest possible quantity
	quantityRangeHigh int64   // highest possible quantity
	buyRangeLow       float64 // lowest possible buy price for resource
	buyRangeHigh      float64 // highest possible buy price for resource
	sellDelta         float64 // highest potential difference of selling price. Selling price is always lower than buying price

	// If set, this function will be called before this item can be bought. Can
	// be used to check for prereqs to purchasing equipment.
	prebuy func(player *Player, item *Item)
}

// Adjust the market price for this listing.
func (l *Listing) adjustMarketPrice() {
	l.Quantity = randInt64(l.quantityRangeLow, l.quantityRangeHigh)
	l.BuyPrice = roundFloat64(randFloat64(l.buyRangeLow, l.buyRangeHigh), 2)
	l.SellPrice = roundFloat64(l.BuyPrice-randFloat64(0.01, l.sellDelta), 2)
}
