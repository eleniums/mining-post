package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/eleniums/mining-post/game"
)

type ListMarketStockResponse struct {
	NextMarketUpdate string          `json:"nextMarketUpdate"`
	Stock            []*game.Listing `json:"stock"`
}

// List entire market inventory.
func (s *Server) ListMarketStock(w http.ResponseWriter, req *http.Request) {
	listings := s.manager.GetMarketStock()
	nextUpdate := s.manager.NextUpdate.Format(time.RFC3339)

	// apply filters if requested
	filterParam := req.URL.Query().Get("filter")
	if filterParam != "" {
		filters := strings.Split(filterParam, ",")
		for _, filter := range filters {
			split := strings.Split(filter, "=")
			if len(split) != 2 {
				http.Error(w, "filter must be a comma-separated list of filters in the format: property=value", http.StatusBadRequest)
				return
			}
			filterType := split[0]
			filterValue := split[1]
			switch filterType {
			case FILTER_NAME:
				listings = game.Filter(listings, func(val *game.Listing) bool {
					return val.Name == filterValue
				})
			case FILTER_TYPE:
				listings = game.Filter(listings, func(val *game.Listing) bool {
					return string(val.Type) == filterValue
				})
			}
		}
	}

	resp := ListMarketStockResponse{
		NextMarketUpdate: nextUpdate,
		Stock:            listings,
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

// Buy an item from the market.
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

type SellOrderRequest struct {
	PlayerName string `json:"player"`
	ItemName   string `json:"item"`
	Quantity   int64  `json:"quantity"`
}

type SellOrderResponse struct {
	Profit  float64 `json:"profit"`
	Message string  `json:"message"`
}

// Sell an item on the market.
func (s *Server) SellOrder(w http.ResponseWriter, req *http.Request) {
	var in SellOrderRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profit, err := s.manager.SellOrder(in.PlayerName, in.ItemName, in.Quantity)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to sell %d of item: %s, err: %v", in.Quantity, in.ItemName, err), http.StatusBadRequest)
		return
	}

	resp := SellOrderResponse{
		Profit:  profit,
		Message: fmt.Sprintf("Successfully sold %d of item: %s, total profit: %.2f", in.Quantity, in.ItemName, profit),
	}

	writeResponse(w, resp)
}
