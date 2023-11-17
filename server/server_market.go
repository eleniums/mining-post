package server

import (
	"net/http"

	"github.com/eleniums/mining-post/models"
)

// List stats for player, including the entire inventory.
func (s *Server) MarketListStock(w http.ResponseWriter, req *http.Request) {
	var in models.MarketListStockRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO: list market stock
	resp := models.MarketListStockResponse{}

	// id, err := s.items.Add(mapItemToDBItem(&in))
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("error saving item: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	writeResponse(w, resp)
}
