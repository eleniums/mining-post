package game

type Market struct {
	Stock []Listing `json:"stock"`
}

type Listing struct {
	Resource

	Quantity  int64   `json:"quantity"`
	SellPrice float64 `json:"sell_price"`
	BuyPrice  float64 `json:"buy_price"`

	buyRangeLow  float64 // lowest possible buy price for resource
	buyRangeHigh float64 // highest possible buy price for resource
	sellDelta    float64 // highest potential difference of selling price. Selling price is always lower than buying price
}
