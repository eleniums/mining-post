package server

import (
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/game"
)

type ListMarketStockResponse struct {
	Stock []*game.Listing `json:"stock"`
}

// List entire market inventory.
func (s *Server) ListMarketStock(w http.ResponseWriter, req *http.Request) {
	listings := s.manager.GetMarketStock()

	resp := ListMarketStockResponse{
		Stock: listings,
	}

	writeResponse(w, resp)
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

// List entire market inventory.
func (s *Server) BuyOrder(w http.ResponseWriter, req *http.Request) {
	var in BuyOrderRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := s.manager.BuyOrder(in.PlayerName, in.ItemName, in.Quantity)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to purchase %d of item: %s, err: %v", in.Quantity, in.ItemName, err), http.StatusBadRequest)
		return
	}

	resp := BuyOrderResponse{
		Cost:    cost,
		Message: fmt.Sprintf("Successfully purchased %d of item: %s, total cost: %.2f", in.Quantity, in.ItemName, cost),
	}

	writeResponse(w, resp)
}

// List entire market inventory.
func (s *Server) SellOrder(w http.ResponseWriter, req *http.Request) {
	// s.manager.SellOrder()

	// resp := ListMarketStockResponse{
	// 	Stock: listings,
	// }

	// writeResponse(w, resp)
}
