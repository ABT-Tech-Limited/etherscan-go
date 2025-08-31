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

// Module: Contract

type ContractABIResp BaseResp

func (r *ContractABIResp) GetData() (string, error) {
	var msg string
	if r.Status == 0 {
		// 尝试解析为字符串（错误信息）
		if err := json.Unmarshal(r.Result, &msg); err == nil {
			return "", fmt.Errorf("API call error: %s", msg)
		}
		return "", fmt.Errorf("API call error with unknown message: %v", r.Result)
	}
	// 尝试解析为ABI字符串
	if err := json.Unmarshal(r.Result, &msg); err != nil {
		return "", fmt.Errorf("failed to parse result: %v", err)
	}
	return msg, nil
}

type ContractSourcecodeResp BaseResp

func (r *ContractSourcecodeResp) GetData() ([]ContractSourceCode, error) {
	if r.Status == 0 {
		return nil, fmt.Errorf("API call error: %s", r.Result)
	}
	var codes []ContractSourceCode
	err := json.Unmarshal(r.Result, &codes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return codes, nil
}

type ContractSourceCode struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	CompilerType         string `json:"CompilerType"`
	OptimizationUsed     string `json:"OptimizationUsed"`
	Runs                 string `json:"Runs"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	ContractFileName     string `json:"ContractFileName"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
	SimilarMatch         string `json:"SimilarMatch"`
}

type ContractCreatorTxInfoResp BaseResp

type ContractCreatorTxInfo struct {
	ContractAddress  string `json:"contractAddress"`
	ContractCreator  string `json:"contractCreator"`
	TxHash           string `json:"txHash"`
	BlockNumber      string `json:"blockNumber"`
	Timestamp        string `json:"timeStamp"`
	ContractFactory  string `json:"contractFactory"`
	CreationBytecode string `json:"creationBytecode"`
}

func (r *ContractCreatorTxInfoResp) GetData() ([]ContractCreatorTxInfo, error) {
	if r.Status == 0 {
		return nil, fmt.Errorf("API call error: %s", r.Result)
	}
	var infos []ContractCreatorTxInfo
	err := json.Unmarshal(r.Result, &infos)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return infos, nil
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
func (r *LogResp) GetData() ([]Log, error) {
	// 如果状态不是成功且不是"没有记录"的情况，检查result是否是错误字符串
	if r.Status == 0 && r.Message != "No records found" {
		// 尝试解析为字符串（错误信息）
		var errorMsg string
		if err := json.Unmarshal(r.Result, &errorMsg); err == nil {
			return nil, fmt.Errorf("API call error: %s", errorMsg)
		}
	}
	// 如果是"No records found"，直接返回空数组
	if r.Status == 0 && r.Message == "No records found" {
		return []Log{}, nil
	}
	// 尝试解析为日志数组
	var logs []Log
	if err := json.Unmarshal(r.Result, &logs); err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return logs, nil
}
