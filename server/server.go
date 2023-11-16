package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/eleniums/mining-post/mem"
	"github.com/go-chi/chi"
)

// ItemStorage defines actions for storing and retrieving items.
type ItemStorage interface {
	Add(item *mem.Item) (string, error)
	Update(id string, item *mem.Item) error
	Get(id string) (*mem.Item, error)
	GetAll() (map[string]*mem.Item, error)
	Remove(id string) error
}

// Server contains the implementation.
type Server struct {
	items ItemStorage
}

// NewServer creates a new instance of Server.
func NewServer(items ItemStorage) *Server {
	return &Server{
		items: items,
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

// getURLParamItemID will return the itemID URL parameter or an empty string if it doesn't exist.
func getURLParamItemID(r *http.Request) string {
	return chi.URLParam(r, "itemID")
}
