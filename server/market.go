package server

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/eleniums/mining-post/game"
)

type ListMarketStockResponse struct {
	NextMarketUpdate string    `json:"nextMarketUpdate"`
	Stock            []Listing `json:"stock"`
}

// List entire market inventory.
func (s *Server) ListMarketStock(w http.ResponseWriter, req *http.Request) {
	// apply filters if requested
	listingFilters := []game.ListingFilter{}
	filterParam := req.URL.Query().Get("filter")
	if filterParam != "" {
		filters := strings.Split(filterParam, ",")
		for _, filter := range filters {
			split := strings.Split(filter, "=")
			if len(split) != 2 {
				http.Error(w, "filter must be a comma-separated list of filters in the format: property=value", http.StatusBadRequest)
				return
			}
			filterProperty := split[0]
			filterValue := split[1]
			listingFilters = append(listingFilters, game.ListingFilter{
				Property: filterProperty,
				Value:    filterValue,
			})
		}
	}

	listings := s.manager.GetMarketStock(listingFilters...)
	nextUpdate := s.manager.NextUpdate.Format(time.RFC3339)

	stock := make([]Listing, len(listings))
	for i, v := range listings {
		stock[i] = NewListing(v)
	}

	resp := ListMarketStockResponse{
		NextMarketUpdate: nextUpdate,
		Stock:            stock,
	}

	writeResponse(w, resp)
}

type BuyOrderRequest struct {
	PlayerName   string `json:"player"`
	ResourceName string `json:"resource"`
	Quantity     int64  `json:"quantity"`
}

type BuyOrderResponse struct {
	Cost    string `json:"cost"`
	Message string `json:"message"`
}

// Buy an item from the market.
func (s *Server) BuyOrder(w http.ResponseWriter, req *http.Request) {
	var in BuyOrderRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cost, err := s.manager.BuyOrder(in.PlayerName, in.ResourceName, in.Quantity)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to purchase %d of resource: %s, err: %v", in.Quantity, in.ResourceName, err), http.StatusBadRequest)
		return
	}

	resp := BuyOrderResponse{
		Cost:    fmt.Sprintf("$%.2f", cost),
		Message: fmt.Sprintf("Successfully purchased %d of resource: %s, total cost: $%.2f", in.Quantity, in.ResourceName, cost),
	}

	writeResponse(w, resp)
}

type SellOrderRequest struct {
	PlayerName   string `json:"player"`
	ResourceName string `json:"resource"`
	Quantity     int64  `json:"quantity"`
}

type SellOrderResponse struct {
	Profit  string `json:"profit"`
	Message string `json:"message"`
}

// Sell an item on the market.
func (s *Server) SellOrder(w http.ResponseWriter, req *http.Request) {
	var in SellOrderRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	profit, err := s.manager.SellOrder(in.PlayerName, in.ResourceName, in.Quantity)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to sell %d of resource: %s, err: %v", in.Quantity, in.ResourceName, err), http.StatusBadRequest)
		return
	}

	resp := SellOrderResponse{
		Profit:  fmt.Sprintf("$%.2f", profit),
		Message: fmt.Sprintf("Successfully sold %d of resource: %s, total profit: $%.2f", in.Quantity, in.ResourceName, profit),
	}

	writeResponse(w, resp)
}
