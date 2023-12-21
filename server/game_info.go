package server

import (
	"net/http"

	"github.com/eleniums/mining-post/game"
	"github.com/eleniums/mining-post/models"
)

// List entire market inventory.
func (s *Server) GameInfo(w http.ResponseWriter, req *http.Request) {
	resp := models.GameInfoResponse{
		Info: game.InfoFile,
	}

	writeResponse(w, resp)
}
