package server

import (
	"net/http"
)

type DigActionRequest struct {
	PlayerName string `json:"player"`
	ItemName   string `json:"item"`
	Quantity   int64  `json:"quantity"`
}

type DigActionResponse struct {
	Cost    float64 `json:"cost"`
	Message string  `json:"message"`
}

// Buy an item from the market.
func (s *Server) DigAction(w http.ResponseWriter, req *http.Request) {
	// TODO: dig action
	// var in DigActionRequest
	// err := readBody(req, &in)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// cost, err := s.manager.DigAction(in.PlayerName, in.ItemName, in.Quantity)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to purchase %d of item: %s, err: %v", in.Quantity, in.ItemName, err), http.StatusBadRequest)
	// 	return
	// }

	// resp := DigActionResponse{
	// 	Cost:    cost,
	// 	Message: fmt.Sprintf("Successfully purchased %d of item: %s, total cost: $%.2f", in.Quantity, in.ItemName, cost),
	// }

	// writeResponse(w, resp)
}

type ProspectActionRequest struct {
	PlayerName string `json:"player"`
	ItemName   string `json:"item"`
	Quantity   int64  `json:"quantity"`
}

type ProspectActionResponse struct {
	Profit  float64 `json:"profit"`
	Message string  `json:"message"`
}

// Sell an item on the market.
func (s *Server) ProspectAction(w http.ResponseWriter, req *http.Request) {
	// TODO: prospect action
	// var in ProspectActionRequest
	// err := readBody(req, &in)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// profit, err := s.manager.ProspectAction(in.PlayerName, in.ItemName, in.Quantity)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to sell %d of item: %s, err: %v", in.Quantity, in.ItemName, err), http.StatusBadRequest)
	// 	return
	// }

	// resp := ProspectActionResponse{
	// 	Profit:  profit,
	// 	Message: fmt.Sprintf("Successfully sold %d of item: %s, total profit: $%.2f", in.Quantity, in.ItemName, profit),
	// }

	// writeResponse(w, resp)
}
