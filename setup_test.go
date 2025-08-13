package etherscan

import (
	"os"
	"time"
)

var api *Client
var chainIdARB = uint64(42161)

func init() {
	apiKey := os.Getenv("EtherscanAPIKey")
	api = New(apiKey, Options{
		Timeout:       time.Second * 10,
		Verbose:       false,
		BeforeRequest: FreeRateLimiter(),
	})
}
