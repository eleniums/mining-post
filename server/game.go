package server

import (
	"net/http"

	"github.com/eleniums/mining-post/game"
)

type GameInfoResponse struct {
	Info string `json:"info"`
}

// Retrieve information about the game.
func (s *Server) GameInfo(w http.ResponseWriter, req *http.Request) {
	resp := GameInfoResponse{
		Info: game.InfoFile,
	}

	writeResponse(w, resp)
}
