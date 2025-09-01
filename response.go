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
// }
//
// {
//   "status": "1",}
//   "message": "OK",
//   "result": [ ... ]
// }
//
//

type BaseResp struct {
	Status  int             `json:"status,string"` // 1 for good, 0 for error
	Message string          `json:"message"`       // OK for good, other words when Status equals 0
	Result  json.RawMessage `json:"result"`
}

type StringResp struct {
	Status  int    `json:"status,string"`
	Message string `json:"message"`
	Result  string `json:"result"`
}

// Module: Contract

type ContractSourcecodeResp BaseResp

func (r *ContractSourcecodeResp) GetData() ([]ContractSourceCode, error) {
	if r.Status == 0 {
		return nil, fmt.Errorf("API call error: %s", r.Result)
	}
	var resultStr string
	if err := json.Unmarshal(r.Result, &resultStr); err == nil {
		return nil, fmt.Errorf("API call error: %s", resultStr)
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
	var resultStr string
	if err := json.Unmarshal(r.Result, &resultStr); err == nil {
		return nil, fmt.Errorf("API call error: %s", resultStr)
	}

	var infos []ContractCreatorTxInfo
	err := json.Unmarshal(r.Result, &infos)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return infos, nil
}

type VerifySourceCodeResp BaseResp

func (r *VerifySourceCodeResp) GetData() (string, error) {
	if r.Status == 0 {
		return "", fmt.Errorf("API call error: %s", r.Result)
	}
	var resultStr string
	if err := json.Unmarshal(r.Result, &resultStr); err != nil {
		return "", fmt.Errorf("failed to parse result: %v", err)
	}
	if resultStr == "Contract source code already verified" {
		return "", nil
	}
	return resultStr, nil
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
	if r.Status == 0 {
		return nil, fmt.Errorf("API call error: %s", r.Result)
	}
	var resultStr string
	if err := json.Unmarshal(r.Result, &resultStr); err == nil {
		return nil, fmt.Errorf("API call error: %s", resultStr)
	}

	var logs []Log
	if err := json.Unmarshal(r.Result, &logs); err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return logs, nil
}
