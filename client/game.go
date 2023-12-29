package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/eleniums/mining-post/game"
	"github.com/eleniums/mining-post/server"
)

type GameClient struct {
	rootURL string
	client  *HTTPClient
}

func NewGameClient(url string, c ...Config) *GameClient {
	return &GameClient{
		rootURL: url,
		client:  NewHTTPClient(c...),
	}
}

func (c *GameClient) GameInfo() (*server.GameInfoResponse, error) {
	return handleClientResp[server.GameInfoResponse](c.client.Get(fmt.Sprintf("%s/game/info", c.rootURL)))
}

func (c *GameClient) GetPlayerInventory(name string) (*server.GetPlayerInventoryResponse, error) {
	return handleClientResp[server.GetPlayerInventoryResponse](c.client.Get(fmt.Sprintf("%s/player/%s/inventory", c.rootURL, name)))

}

func (c *GameClient) ListMarketStock(filters ...game.ListingFilter) (*server.ListMarketStockResponse, error) {
	queryParams := []string{}

	filterParam := []string{}
	for _, v := range filters {
		filterParam = append(filterParam, fmt.Sprintf("%s=%s", v.Property, v.Value))
	}
	if len(filterParam) > 0 {
		queryParams = append(queryParams, "filter", strings.Join(filterParam, ","))
	}

	return handleClientResp[server.ListMarketStockResponse](c.client.Get(fmt.Sprintf("%s/market/stock", c.rootURL), queryParams...))
}

func (c *GameClient) BuyOrder(req server.BuyOrderRequest) (*server.BuyOrderResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	return handleClientResp[server.BuyOrderResponse](c.client.Post(fmt.Sprintf("%s/market/buy", c.rootURL), payload))
}

func (c *GameClient) SellOrder(req server.SellOrderRequest) (*server.SellOrderResponse, error) {
	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	return handleClientResp[server.SellOrderResponse](c.client.Post(fmt.Sprintf("%s/market/sell", c.rootURL), payload))
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
