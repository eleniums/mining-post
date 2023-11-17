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

func (c *GameClient) GetPlayerInventory(name string) (models.PlayerListInventoryResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/player/inventory", c.rootURL))
	if err != nil {
		return models.PlayerListInventoryResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return models.PlayerListInventoryResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp models.PlayerListInventoryResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return models.PlayerListInventoryResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}

func (c *GameClient) ListMarketStock() (models.MarketListStockResponse, error) {
	code, body, err := c.client.Get(fmt.Sprintf("%s/market/stock", c.rootURL))
	if err != nil {
		return models.MarketListStockResponse{}, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return models.MarketListStockResponse{}, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp models.MarketListStockResponse
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return models.MarketListStockResponse{}, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return resp, nil
}
