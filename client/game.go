package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eleniums/mining-post/models"
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

func (c *GameClient) GetPlayerInventory(name string) (models.GetPlayerInventoryResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/player/%s/inventory", c.rootURL, name))
	if err != nil {
		return models.GetPlayerInventoryResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return models.GetPlayerInventoryResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp models.GetPlayerInventoryResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return models.GetPlayerInventoryResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}

func (c *GameClient) ListMarketStock() (models.ListMarketStockResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/market/stock", c.rootURL))
	if err != nil {
		return models.ListMarketStockResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return models.ListMarketStockResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp models.ListMarketStockResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return models.ListMarketStockResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}
