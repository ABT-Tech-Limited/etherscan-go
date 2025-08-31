package etherscan

var logModuleParams = map[string]string{
	"module": "logs",
	"action": "getLogs",
}

// GetEventLogsByAddress Returns the event logs from an address, with optional filtering by block range.
func (c *client) GetEventLogsByAddress(req GetEventLogsByAddressReq) (resp *LogResp, err error) {
	params := logModuleParams
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// GetEventLogsByTopics Returns the events log in a block range, filtered by topics.
func (c *client) GetEventLogsByTopics(req GetEventLogsByTopicsReq) (resp *LogResp, err error) {
	params := logModuleParams
	for k, v := range StructToMap(req) {
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

// GetEventLogsByAddressFilterByTopics Returns the event logs from an address, filtered by topics and block range.
func (c *client) GetEventLogsByAddressFilterByTopics(req GetEventLogsByAddressFilterByTopicsReq) (resp *LogResp, err error) {
	params := logModuleParams
	for k, v := range StructToMap(req) {
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
