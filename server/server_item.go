package server

import (
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/mem"
	"github.com/eleniums/mining-post/models"
)

// InsertItem inserts a new item.
func (s *Server) InsertItem(w http.ResponseWriter, req *http.Request) {
	var in models.Item
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := s.items.Add(mapItemToDBItem(&in))
	if err != nil {
		http.Error(w, fmt.Sprintf("error saving item: %v", err), http.StatusInternalServerError)
		return
	}

	resp := in
	resp.ID = id

	writeResponse(w, resp)
}

// GetItemByID gets an item using an ID.
func (s *Server) GetItemByID(w http.ResponseWriter, req *http.Request) {
	id := getURLParamItemID(req)
	if id == "" {
		http.Error(w, "missing item id", http.StatusBadRequest)
		return
	}

	item, err := s.items.Get(id)
	if err != nil {
		if err == mem.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, fmt.Sprintf("error retrieving item: %v", err), http.StatusInternalServerError)
		return
	}

	resp := mapDBItemToItem(id, item)

	writeResponse(w, resp)
}

// GetItems retrieves all items.
func (s *Server) GetItems(w http.ResponseWriter, req *http.Request) {
	items, err := s.items.GetAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("error retrieving items: %v", err), http.StatusInternalServerError)
		return
	}

	resp := models.Items{
		Items: []*models.Item{},
	}

	for k, v := range items {
		resp.Items = append(resp.Items, mapDBItemToItem(k, v))
	}

	writeResponse(w, &resp)
}

// UpdateItem updates an item.
func (s *Server) UpdateItem(w http.ResponseWriter, req *http.Request) {
	id := getURLParamItemID(req)
	if id == "" {
		http.Error(w, "missing item id", http.StatusBadRequest)
		return
	}

	var in models.Item
	err := readBody(req, &in)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	item := mapItemToDBItem(&in)

	err = s.items.Update(id, item)
	if err != nil {
		if err == mem.ErrNotFound {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, fmt.Sprintf("error updating item: %v", err), http.StatusInternalServerError)
		return
	}
}

// DeleteItem deletes an item.
func (s *Server) DeleteItem(w http.ResponseWriter, req *http.Request) {
	id := getURLParamItemID(req)
	if id == "" {
		http.Error(w, "missing item id", http.StatusBadRequest)
		return
	}

	err := s.items.Remove(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error removing item: %v", err), http.StatusInternalServerError)
		return
	}
}
