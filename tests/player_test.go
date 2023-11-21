package tests

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_GetPlayerInventory_Success(t *testing.T) {
	// act
	resp, err := client.GetPlayerInventory(testPlayer)

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
