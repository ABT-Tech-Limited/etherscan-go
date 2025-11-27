package etherscan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetGasOracle(t *testing.T) {
	rsp, err := api.GetGasOracle(context.Background(), chainIDETH)
	assert.Nil(t, err)
	data, err := rsp.GetData()
	assert.Nil(t, err)
	assert.NotNil(t, data)
}
