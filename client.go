package etherscan

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/time/rate"
	"resty.dev/v3"
)

// Client is the interface for interacting with Etherscan API
type Client interface {
	GetContractABI(req GetContractABIReq) (*ContractABIResp, error)
	GetContractSourceCode(req GetContractSourceCodeReq) (*ContractSourcecodeResp, error)
	GetContractCreatorTxInfo(req GetContractCreatorTxInfoReq) (*ContractCreatorTxInfoResp, error)

	GetEventLogsByAddress(req GetEventLogsByAddressReq) (*LogResp, error)
	GetEventLogsByTopics(req GetEventLogsByTopicsReq) (*LogResp, error)
	GetEventLogsByAddressFilterByTopics(req GetEventLogsByAddressFilterByTopicsReq) (*LogResp, error)

	Debug() Client
}

// client is the concrete implementation of the Client interface
type client struct {
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

func New(apiKey string, opts ...Options) Client {
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

	return &client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

func NewWithClient(apiKey string, restyCli *resty.Client) Client {
	restyCli.AddRequestMiddleware(func(client *resty.Client, request *resty.Request) error {
		request.SetQueryParam("apiKey", apiKey)
		return nil
	})

	return &client{
		resty:  restyCli,
		apiKey: apiKey,
	}
}

func (c *client) Debug() Client {
	c.resty.SetDebug(true)
	return c
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
