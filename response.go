package etherscan

import (
	"encoding/json"
	"fmt"
)

//
// {
//   "status": "0",
//   "message": "NOTOK",
//   "result": "Max calls per sec rate limit reached (5/sec)"
// }
//
// {
//   "status": "0",
//   "message": "No records found",
//   "result": []
//}
//

type BaseResp struct {
	Status  int             `json:"status,string"` // 1 for good, 0 for error
	Message string          `json:"message"`       // OK for good, other words when Status equals 0
	Result  json.RawMessage `json:"result"`
}

// Module: Log

type LogResp BaseResp

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

// GetData 获取日志数据，如果result是错误信息则返回错误
func (lr *LogResp) GetData() ([]Log, error) {
	// 如果状态不是成功且不是"没有记录"的情况，检查result是否是错误字符串
	if lr.Status == 0 && lr.Message != "No records found" {
		// 尝试解析为字符串（错误信息）
		var errorMsg string
		if err := json.Unmarshal(lr.Result, &errorMsg); err == nil {
			return nil, fmt.Errorf("API call error: %s", errorMsg)
		}
	}
	// 如果是"No records found"，直接返回空数组
	if lr.Status == 0 && lr.Message == "No records found" {
		return []Log{}, nil
	}
	// 尝试解析为日志数组
	var logs []Log
	if err := json.Unmarshal(lr.Result, &logs); err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return logs, nil
}
