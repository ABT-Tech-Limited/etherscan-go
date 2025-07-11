package etherscan

import "strconv"

// GetEventLogsByAddress Returns the event logs from an address, with optional filtering by block range.
func (c *Client) GetEventLogsByAddress(chainId uint64, req GetEventLogsByAddressReq) (resp *LogResp, err error) {
	params := map[string]string{
		"apiKey":  c.apiKey,
		"chainid": strconv.FormatUint(chainId, 10),
		"module":  "logs",
		"action":  "getLogs",
	}
	for k, v := range ToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// GetEventLogsByAddressFilterByTopics Returns the event logs from an address, filtered by topics and block range.
func (c *Client) GetEventLogsByAddressFilterByTopics(chainId uint64, req GetEventLogsByAddressFilterByTopicsReq) (resp *LogResp, err error) {
	params := map[string]string{
		"apiKey":  c.apiKey,
		"chainid": strconv.FormatUint(chainId, 10),
		"module":  "logs",
		"action":  "getLogs",
	}
	for k, v := range ToMap(req) {
		if k != "topics" {
			params[k] = v
		}
	}
	for k, v := range req.Topics {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}
