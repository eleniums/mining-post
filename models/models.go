package models

import (
	"github.com/eleniums/mining-post/game"
)

type GetPlayerInventoryResponse struct {
	Player *game.Player `json:"player"`
}

type ListMarketStockResponse struct {
	NextMarketUpdate string          `json:"nextMarketUpdate"`
	Stock            []*game.Listing `json:"stock"`
}

type BuyOrderRequest struct {
	PlayerName string `json:"player"`
	ItemName   string `json:"item"`
	Quantity   int64  `json:"quantity"`
}

type BuyOrderResponse struct {
	Cost    float64 `json:"cost"`
	Message string  `json:"message"`
}

type SellOrderRequest struct {
	PlayerName string `json:"player"`
	ItemName   string `json:"item"`
	Quantity   int64  `json:"quantity"`
}

type SellOrderResponse struct {
	Profit  float64 `json:"profit"`
	Message string  `json:"message"`
}

type GameInfoResponse struct {
	Info string `json:"info"`
}
