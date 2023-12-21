package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/eleniums/mining-post/client"
	"github.com/eleniums/mining-post/models"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_ListMarketStock_Success(t *testing.T) {
	// act
	resp, err := gameClient.ListMarketStock()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Greater(t, len(resp.Stock), 1)

	nextUpdate, _ := time.Parse(time.RFC3339, resp.NextMarketUpdate)
	nextUpdate = nextUpdate.Local()
	fmt.Printf("Next market update is at: %v\n", nextUpdate.Format(time.RFC1123))
}

func Test_Integration_ListMarketStock_Filtered(t *testing.T) {
	// act
	resp, err := gameClient.ListMarketStock(
		client.Filter{
			Property: "Type",
			Value:    "Commodity",
		},
		client.Filter{
			Property: "Name",
			Value:    "Limestone",
		},
	)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.Len(t, resp.Stock, 1)

	nextUpdate, _ := time.Parse(time.RFC3339, resp.NextMarketUpdate)
	nextUpdate = nextUpdate.Local()
	fmt.Printf("Next market update is at: %v\n", nextUpdate.Format(time.RFC1123))
}

func Test_Integration_BuyOrder_Success(t *testing.T) {
	req := models.BuyOrderRequest{
		PlayerName: "snelson",
		ItemName:   "Limestone",
		Quantity:   3,
	}

	// act
	resp, err := gameClient.BuyOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Integration_SellOrder_Success(t *testing.T) {
	req := models.SellOrderRequest{
		PlayerName: "snelson",
		ItemName:   "Limestone",
		Quantity:   3,
	}

	// act
	resp, err := gameClient.SellOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
