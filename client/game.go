package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/eleniums/mining-post/models"
)

type Filter struct {
	Property string
	Value    string
}

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

func (c *GameClient) GameInfo() (*models.GameInfoResponse, error) {
	return handleClientResp[models.GameInfoResponse](c.client.Get(fmt.Sprintf("%s/game/info", c.rootURL)))
}

func (c *GameClient) GetPlayerInventory(name string) (*models.GetPlayerInventoryResponse, error) {
	return handleClientResp[models.GetPlayerInventoryResponse](c.client.Get(fmt.Sprintf("%s/player/%s/inventory", c.rootURL, name)))

}

func (c *GameClient) ListMarketStock(filters ...Filter) (*models.ListMarketStockResponse, error) {
	queryParams := []string{}

	filterParam := []string{}
	for _, v := range filters {
		filterParam = append(filterParam, fmt.Sprintf("%s=%s", v.Property, v.Value))
	}
	if len(filterParam) > 0 {
		queryParams = append(queryParams, "filter", strings.Join(filterParam, ","))
	}

	return handleClientResp[models.ListMarketStockResponse](c.client.Get(fmt.Sprintf("%s/market/stock", c.rootURL), queryParams...))
}

func (c *GameClient) BuyOrder(req models.BuyOrderRequest) (*models.BuyOrderResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	return handleClientResp[models.BuyOrderResponse](c.client.Post(fmt.Sprintf("%s/market/buy", c.rootURL), payload))
}

func (c *GameClient) SellOrder(req models.SellOrderRequest) (*models.SellOrderResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	return handleClientResp[models.SellOrderResponse](c.client.Post(fmt.Sprintf("%s/market/sell", c.rootURL), payload))
}

func handleClientResp[T any](code int, body []byte, err error) (*T, error) {
	if err != nil {
		return nil, fmt.Errorf("error calling service: %v", err)
	}
	if code != http.StatusOK {
		return nil, fmt.Errorf("invalid status code: %d, message: %v", code, string(body))
	}

	var resp T
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response body: %v", err)
	}

	return &resp, nil
}
