package server

import (
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/models"
)

// List stats for player, including the entire inventory.
func (s *Server) GetPlayerInventory(w http.ResponseWriter, req *http.Request) {
	var in models.GetPlayerInventoryRequest
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	player, err := s.manager.GetPlayer(in.Name)
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving player: %v", err), http.StatusBadRequest)
		return
	}

	resp := models.GetPlayerInventoryResponse{
		Player: player,
	}

	writeResponse(w, resp)
}
