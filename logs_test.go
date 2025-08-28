package etherscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TEST TX
// ARB: https://arbiscan.io/tx/0xa29d163e1382e6f879e97e89af12a400f1384561946beba7a1f9ee49a5192579

func TestClient_GetEventLogsByAddress(t *testing.T) {
	page := 1
	offset := 0
	rsp, err := api.GetEventLogsByAddress(GetEventLogsByAddressReq{
		ChainID:   chainIdARB,
		Address:   "0x56785734c9DC26BBc6B8AA9a615Fa19eFb0b677b",
		FromBlock: 356174872,
		ToBlock:   356174872,
		Page:      &page,
		Offset:    &offset,
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	logs, err := rsp.GetData()
	assert.Equal(t, 3, len(logs))
	assert.Nil(t, err)
}

func TestClient_GetEventLogsByTopics(t *testing.T) {
	page := 1
	offset := 0
	rsp, err := api.GetEventLogsByTopics(GetEventLogsByTopicsReq{
		ChainID:   chainIdARB,
		FromBlock: 356174872,
		ToBlock:   356174872,
		Topics: map[string]string{
			"topic0": "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0",
		},
		Page:   &page,
		Offset: &offset,
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	logs, err := rsp.GetData()
	assert.Equal(t, 1, len(logs))
	assert.Nil(t, err)
}

func TestClient_GetEventLogsByAddressFilterByTopics(t *testing.T) {
	page := 1
	offset := 0
	rsp, err := api.GetEventLogsByAddressFilterByTopics(GetEventLogsByAddressFilterByTopicsReq{
		ChainID:   chainIdARB,
		Address:   "0x56785734c9DC26BBc6B8AA9a615Fa19eFb0b677b",
		FromBlock: 356174872,
		ToBlock:   356174872,
		Topics: map[string]string{
			"topic0":       "0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0",
			"topic1":       "0x0000000000000000000000000000000000000000000000000000000000000000",
			"topic2":       "0x000000000000000000000000cc6a237534e983f3e9c2074d49f2b9a112101ab1",
			"topic1_2_opr": "and",
		},
		Page:   &page,
		Offset: &offset,
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	logs, err := rsp.GetData()
	assert.Equal(t, 1, len(logs))
	assert.Nil(t, err)
}
