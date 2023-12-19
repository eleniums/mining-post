package tests

import (
	"fmt"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_Integration_GameInfo_Success(t *testing.T) {
	// act
	resp, err := gameClient.GameInfo()

	// assert
	assert.NoError(t, err)
	assert.NotNil(t, resp)

	assert.NotEmpty(t, resp.Info)

	fmt.Printf("\n%s\n", resp.Info)
}
