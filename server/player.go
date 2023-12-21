package server

import (
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/models"
	"github.com/go-chi/chi"
)

// List stats for player, including the entire inventory.
func (s *Server) GetPlayerInventory(w http.ResponseWriter, req *http.Request) {
	playerName := chi.URLParam(req, "player-name")

	player, err := s.manager.GetPlayer(playerName)
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving player: %v", err), http.StatusBadRequest)
		return
	}

	resp := models.GetPlayerInventoryResponse{
		Player: player,
	}

	writeResponse(w, resp)
}
