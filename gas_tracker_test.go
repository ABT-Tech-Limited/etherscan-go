package etherscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetGasOracle(t *testing.T) {
	rsp, err := api.GetGasOracle(chainIDETH)
	assert.Nil(t, err)
	data, err := rsp.GetData()
	assert.Nil(t, err)
	assert.NotNil(t, data)
}
