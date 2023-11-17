package server

import (
	"net/http"

	"github.com/eleniums/mining-post/models"
)

// List entire market inventory.
func (s *Server) ListMarketStock(w http.ResponseWriter, req *http.Request) {
	market := s.manager.GetMarketStock()

	resp := models.ListMarketStockResponse{
		Market: market,
	}

	writeResponse(w, resp)
}
