package etherscan

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetEventLogsByAddress(t *testing.T) {
	rsp, err := api.GetEventLogsByAddress(chainIdARB, GetEventLogsByAddressReq{
		Address:   "0x56785734c9DC26BBc6B8AA9a615Fa19eFb0b677b",
		FromBlock: 356426141,
		ToBlock:   356427141,
		Page:      1,
		Offset:    0,
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	assert.Equal(t, 10, len(rsp.Result))
}

func TestClient_GetEventLogsByTopics(t *testing.T) {
	rsp, err := api.GetEventLogsByTopics(chainIdARB, GetEventLogsByTopicsReq{
		FromBlock: 356426141,
		ToBlock:   356427141,
		Page:      1,
		Offset:    0,
		Topics: map[string]string{
			"topic0": "0x2a3de20682fb291f444b5c1469d7e0950c558ce3dadf97163687873e29bcf4ae",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	assert.Equal(t, 2, len(rsp.Result))
}

func TestClient_GetEventLogsByAddressFilterByTopics(t *testing.T) {
	rsp, err := api.GetEventLogsByAddressFilterByTopics(chainIdARB, GetEventLogsByAddressFilterByTopicsReq{
		Address:   "0x56785734c9DC26BBc6B8AA9a615Fa19eFb0b677b",
		FromBlock: 356426141,
		ToBlock:   356427141,
		Page:      1,
		Offset:    0,
		Topics: map[string]string{
			"topic0": "0x2a3de20682fb291f444b5c1469d7e0950c558ce3dadf97163687873e29bcf4ae",
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, rsp.Status, 1)
	assert.Equal(t, rsp.Message, "OK")
	assert.Equal(t, 2, len(rsp.Result))
}
