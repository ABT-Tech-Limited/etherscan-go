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

// New tests for omitempty logic
func TestStructToMap_OmitEmpty(t *testing.T) {
	type sample struct {
		A string `json:"a,omitempty"`
		B string `json:"b"`
		C int    `json:"c,omitempty"`
		D int    `json:"d,omitempty"`
		E []int  `json:"e,omitempty"`
		F []int  `json:"f"`
		G *int   `json:"g,omitempty"`
		H *int   `json:"h,omitempty"`
	}
	ptrVal := 10
	s := sample{A: "", B: "present", C: 0, D: 5, E: nil, F: []int{}, G: nil, H: &ptrVal}
	m := StructToMap(s)
	// a should be omitted because empty string with omitempty
	_, hasA := m["a"]
	assert.False(t, hasA)
	// b should be present
	assert.Equal(t, "present", m["b"])
	// c omitted (zero int)
	_, hasC := m["c"]
	assert.False(t, hasC)
	// d present (non-zero int)
	assert.Equal(t, "5", m["d"])
	// e omitted (nil slice)
	_, hasE := m["e"]
	assert.False(t, hasE)
	// f present (empty slice but no omitempty tag)
	assert.Equal(t, "[]", m["f"])
	// g omitted (nil pointer)
	_, hasG := m["g"]
	assert.False(t, hasG)
	// h present (non-nil pointer)
	assert.Equal(t, "10", m["h"])
}
