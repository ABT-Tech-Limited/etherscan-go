package etherscan

import (
	"os"
	"time"

	"resty.dev/v3"
)

var api Client
var chainIdARB = uint64(42161)
var chainIDETH = uint64(1)

func init() {
	apiKey := os.Getenv("EtherscanAPIKey")
	api = New(apiKey, Options{
		Timeout:       time.Second * 10,
		Verbose:       false,
		BeforeRequest: []resty.RequestMiddleware{FreeRateLimiter()},
	})
}
