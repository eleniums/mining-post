package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/eleniums/mining-post/game"
	"github.com/eleniums/mining-post/server"

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
		game.ListingFilter{
			Property: "Type",
			Value:    "Commodity",
		},
		game.ListingFilter{
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
	req := server.BuyOrderRequest{
		PlayerName: testPlayer,
		ItemName:   "Limestone",
		Quantity:   2,
	}

	// act
	resp, err := gameClient.BuyOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Integration_BuyOrder_InvalidQuantity(t *testing.T) {
	req := server.BuyOrderRequest{
		PlayerName: testPlayer,
		ItemName:   "Limestone",
		Quantity:   0,
	}

	// act
	resp, err := gameClient.BuyOrder(req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func Test_Integration_SellOrder_Success(t *testing.T) {
	req := server.SellOrderRequest{
		PlayerName: testPlayer,
		ItemName:   "Limestone",
		Quantity:   2,
	}

	// act
	resp, err := gameClient.SellOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Integration_SellOrder_InvalidQuantity(t *testing.T) {
	req := server.SellOrderRequest{
		PlayerName: testPlayer,
		ItemName:   "Limestone",
		Quantity:   0,
	}

	// act
	resp, err := gameClient.SellOrder(req)

	// assert
	assert.Error(t, err)
	assert.Nil(t, resp)
}
