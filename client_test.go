package etherscan

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	mockTimeout := time.Second * 20
	mockBaseUrl := "https://api.etherscan.io/v2/api/test"
	mockVerbose := true
	cli := New("TEST_API_KEY", Options{
		Timeout: mockTimeout,
		BaseUrl: mockBaseUrl,
		Verbose: mockVerbose,
	})
	assert.NotNil(t, cli)
	assert.Equal(t, cli.resty.Timeout(), mockTimeout)
	assert.Equal(t, cli.resty.BaseURL(), mockBaseUrl)
	assert.Equal(t, cli.resty.IsDebug(), mockVerbose)
}
