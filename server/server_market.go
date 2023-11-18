package server

import (
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
