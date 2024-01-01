package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type GetPlayerInventoryResponse struct {
	Player Player `json:"player"`
}

// List stats for player, including their entire inventory.
func (s *Server) GetPlayerInventory(w http.ResponseWriter, req *http.Request) {
	playerName := chi.URLParam(req, "player-name")

	player, err := s.manager.GetPlayer(playerName)
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving player: %v", err), http.StatusBadRequest)
		return
	}

	resp := GetPlayerInventoryResponse{
		Player: NewPlayer(player),
	}

	writeResponse(w, resp)
}
