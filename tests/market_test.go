package tests

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_ListMarketStock_Success(t *testing.T) {
	// act
	resp, err := client.ListMarketStock()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	fmt.Println("resp:", resp)
}
