package etherscan

import (
	"encoding/json"
	"fmt"
)

var logModuleParams = map[string]string{
	"module": "logs",
	"action": "getLogs",
}

// GetEventLogsByAddress Returns the event logs from an address, with optional filtering by block range.
func (c *client) GetEventLogsByAddress(req GetEventLogsByAddressReq) (resp *LogResp, err error) {
	params := CopyMap(logModuleParams)
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// GetEventLogsByTopics Returns the events log in a block range, filtered by topics.
func (c *client) GetEventLogsByTopics(req GetEventLogsByTopicsReq) (resp *LogResp, err error) {
	params := CopyMap(logModuleParams)
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
	params := CopyMap(logModuleParams)
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

// ----------------------------	REQUEST	------------------------------------

// ----------------------------	RESPONSE ------------------------------------

type LogResp BaseResp

// GetData 获取日志数据，如果result是错误信息则返回错误
func (r *LogResp) GetData() ([]Log, error) {
	noData, err := (*BaseResp)(r).Parse()
	if err != nil {
		return nil, err
	}
	if noData {
		return []Log{}, nil
	}

	var logs []Log
	if err = json.Unmarshal(r.Result, &logs); err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return logs, nil
}

type Log struct {
	Address          string   `json:"address"`
	Topics           []string `json:"topics"`
	Data             string   `json:"data"`
	BlockNumber      string   `json:"blockNumber"`
	BlockHash        string   `json:"blockHash"`
	TimeStamp        string   `json:"timeStamp"`
	GasPrice         string   `json:"gasPrice"`
	GasUsed          string   `json:"gasUsed"`
	LogIndex         string   `json:"logIndex"`
	TransactionHash  string   `json:"transactionHash"`
	TransactionIndex string   `json:"transactionIndex"`
}
