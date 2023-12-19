package server

import (
	"net/http"
)

// Ping service for availability.
func (s *Server) Ping(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
}
