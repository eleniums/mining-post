package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/server"
)

type GameClient struct {
	rootURL string
	client  *HTTPClient
}

func NewGameClient(url string) *GameClient {
	return &GameClient{
		rootURL: url,
		client:  NewHTTPClient(),
	}
}

func (c *GameClient) GetPlayerInventory(name string) (server.GetPlayerInventoryResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/player/%s/inventory", c.rootURL, name))
	if err != nil {
		return server.GetPlayerInventoryResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return server.GetPlayerInventoryResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp server.GetPlayerInventoryResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return server.GetPlayerInventoryResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}

func (c *GameClient) ListMarketStock() (server.ListMarketStockResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/market/stock", c.rootURL))
	if err != nil {
		return server.ListMarketStockResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return server.ListMarketStockResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp server.ListMarketStockResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return server.ListMarketStockResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}

func (c *GameClient) BuyOrder(req server.BuyOrderRequest) (server.BuyOrderResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return server.BuyOrderResponse{}, fmt.Errorf("error marshaling request: %v", err)
	}

	code, body, err := c.client.Post(fmt.Sprintf("%s/market/buy", c.rootURL), payload)
	if err != nil {
		return server.BuyOrderResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return server.BuyOrderResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp server.BuyOrderResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return server.BuyOrderResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}
