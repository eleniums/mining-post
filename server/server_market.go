package server

import (
	"net/http"

	"github.com/eleniums/mining-post/game"
)

type ListMarketStockResponse struct {
	game.Market
}

// List entire market inventory.
func (s *Server) ListMarketStock(w http.ResponseWriter, req *http.Request) {
	market := s.manager.GetMarketStock()

	resp := ListMarketStockResponse{
		Market: market,
	}

	writeResponse(w, resp)
}
