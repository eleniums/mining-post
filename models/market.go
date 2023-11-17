package models

type ListMarketStockResponse struct {
	Market
}

type Market struct {
	Stock []Listing `json:"stock"`
}

type Listing struct {
	Resource

	Quantity  int64   `json:"quantity"`
	SellPrice float64 `json:"sell_price"`
	BuyPrice  float64 `json:"buy_price"`
}
