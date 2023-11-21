package tests

import (
	"testing"

	"github.com/eleniums/mining-post/server"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_ListMarketStock_Success(t *testing.T) {
	// act
	resp, err := client.ListMarketStock()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Integration_BuyOrder_Success(t *testing.T) {
	req := server.BuyOrderRequest{
		PlayerName: "snelson",
		ItemName:   "Granite",
		Quantity:   10,
	}

	// act
	resp, err := client.BuyOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}

func Test_Integration_SellOrder_Success(t *testing.T) {
	req := server.SellOrderRequest{
		PlayerName: "snelson",
		ItemName:   "Granite",
		Quantity:   10,
	}

	// act
	resp, err := client.SellOrder(req)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
