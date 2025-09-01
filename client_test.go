package etherscan

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cli := New("TEST_API_KEY")
	assert.NotNil(t, cli)

	// Test with options
	mockTimeout := time.Second * 20
	mockBaseUrl := "https://api.etherscan.io/v2/api/test"
	mockVerbose := true
	cliWithOpts := New("TEST_API_KEY", Options{
		Timeout: mockTimeout,
		BaseUrl: mockBaseUrl,
		Verbose: mockVerbose,
	})
	assert.NotNil(t, cliWithOpts)
}
