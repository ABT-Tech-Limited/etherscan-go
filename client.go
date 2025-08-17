package etherscan

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"resty.dev/v3"
)

type Client struct {
	resty  *resty.Client
	apiKey string
}

type Options struct {
	Timeout       time.Duration
	BaseUrl       string
	Verbose       bool
	Transport     *http.Transport
	BeforeRequest []resty.RequestMiddleware // for rate limit
}

func New(apiKey string, opts ...Options) *Client {
	restyCli := resty.New()
	restyCli.SetTimeout(time.Second * 10)
	restyCli.SetBaseURL("https://api.etherscan.io/v2/api")
	restyCli.SetHeader("User-Agent", "etherscan-go-client")
	restyCli.AddRequestMiddleware(func(client *resty.Client, request *resty.Request) error {
		request.SetQueryParam("apiKey", apiKey)
		return nil
	})
	if len(opts) > 0 {
		opt := opts[0]
		restyCli.SetTimeout(opt.Timeout)
		restyCli.SetDebug(opt.Verbose)
		if opt.BaseUrl != "" {
			restyCli.SetBaseURL(opt.BaseUrl)
		}
		if opt.Transport != nil {
			restyCli.SetTransport(opt.Transport)
		}
		for _, f := range opt.BeforeRequest {
			restyCli.AddRequestMiddleware(f)
		}
	}

	return &Client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

func NewWithClient(apiKey string, restyCli *resty.Client) *Client {
	restyCli.AddRequestMiddleware(func(client *resty.Client, request *resty.Request) error {
		request.SetQueryParam("apiKey", apiKey)
		return nil
	})

	return &Client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

func FreeRateLimiter() func(client *resty.Client, req *resty.Request) error {
	limiter := rate.NewLimiter(5, 1)
	return func(client *resty.Client, req *resty.Request) error {
		if err := limiter.Wait(req.Context()); err != nil {
			return fmt.Errorf("free rate limit exceeded: %w", err)
		}
		return nil
	}
}
