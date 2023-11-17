package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eleniums/mining-post/game"
)

// Server contains the implementation.
type Server struct {
	gameManager *game.GameManager
}

// NewServer creates a new instance of Server.
func NewServer(gm *game.GameManager) *Server {
	return &Server{
		gameManager: gm,
	}
}

// readBody will parse the request body into a given struct.
func readBody(req *http.Request, msg interface{}) error {
	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, msg)
}

// writeResponse will convert the response to JSON and write it.
func writeResponse(w http.ResponseWriter, resp interface{}) {
	r, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(r)
}
