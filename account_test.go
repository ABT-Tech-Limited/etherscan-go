package etherscan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetNormalTransactions(t *testing.T) {
	rsp, err := api.GetNormalTransactionsByAddress(context.Background(), GetNormalTransactionsByAddressReq{
		ChainID: chainIDSepolia,
		Address: "0xdb6b74204C641F8A4e10fe1A4A003386f6534447",
		Page:    1,
		Offset:  10,
		Sort:    "desc",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, rsp.Status)
	txs, err := rsp.GetData()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(txs))
}

func TestClient_GetERC20TokenTransferEvents(t *testing.T) {
	rsp, err := api.GetERC20TokenTransferByAddress(context.Background(), GetERC20TokenTransferEventsReq{
		ChainID:         chainIDSepolia,
		ContractAddress: "0xdb6b74204C641F8A4e10fe1A4A003386f6534447",
		Page:            1,
		Offset:          10,
		Sort:            "asc",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, rsp.Status)
	transfers, err := rsp.GetData()
	assert.Nil(t, err)
	assert.Equal(t, 10, len(transfers))
}
