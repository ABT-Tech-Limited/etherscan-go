package etherscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetNormalTransactions(t *testing.T) {
	rsp, err := api.GetNormalTransactions(GetNormalTransactionsReq{
		ChainID: chainIDSepolia,
		Address: "0xdb6b74204C641F8A4e10fe1A4A003386f6534447",
		Page:    1,
		Offset:  10,
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, rsp.Status)
	_, err = rsp.GetData()
	assert.Nil(t, err)
}
