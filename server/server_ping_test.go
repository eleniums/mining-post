package server

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eleniums/mining-post/mem"

	assert "github.com/stretchr/testify/require"
)

func Test_Unit_Server_Ping_Success(t *testing.T) {
	// arrange
	cache := mem.NewCache()
	server := NewServer(cache)

	w := httptest.NewRecorder()
	rq := httptest.NewRequest(http.MethodGet, "/ping", nil)

	// act
	server.Ping(w, rq)

	// assert
	r := w.Result()
	assert.Equal(t, http.StatusOK, r.StatusCode)
}

func unmarshalBody(r *http.Response, v interface{}) error {
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, v)
}
