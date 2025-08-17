package etherscan

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructToMap(t *testing.T) {
	page := 1
	offset := 0
	req := GetEventLogsByAddressReq{
		ChainID:   chainIdARB,
		Address:   "0x56785734c9DC26BBc6B8AA9a615Fa19eFb0b677b",
		FromBlock: 356174872,
		ToBlock:   356174872,
		Page:      &page,
		Offset:    &offset,
	}
	m := StructToMap(req)
	assert.Equal(t, strconv.FormatInt(int64(page), 10), m["page"])
	assert.Equal(t, strconv.FormatInt(int64(offset), 10), m["offset"])
}
