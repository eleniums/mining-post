package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Server_Ping_Success(t *testing.T) {
	// arrange
	server := NewServer(nil)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, "/ping", nil)

	// act
	server.Ping(w, rq)

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)
}
