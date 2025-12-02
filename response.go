package etherscan

import (
	"encoding/json"
	"fmt"
)

//
// Example responses:
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
//   "status": "1",
//   "message": "OK",
//   "result": [ ... ]
// }
//
//

const (
	NoRecordsFound              = "No records found"
	NoDataFound                 = "No data found"
	NoTransactionsFound         = "No transactions found"
	ContractCodeAlreadyVerified = "Contract source code already verified"
)

func IsNoDataFound(msg string) bool {
	return msg == NoRecordsFound || msg == NoDataFound || msg == NoTransactionsFound
}

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

// Parse 解析基础响应，检查状态和消息
func (r *BaseResp) Parse() (bool, error) {
	if r == nil {
		return false, fmt.Errorf("response is nil")
	}
	if IsNoDataFound(r.Message) { // 没有找到数据
		return true, nil
	}
	if r.Status == 0 {
		// NOT OK 解析Result
		var resultStr string
		if err := json.Unmarshal(r.Result, &resultStr); err == nil {
			return false, fmt.Errorf("status not ok: %s", resultStr)
		}
		return false, fmt.Errorf("status not ok: %s", r.Message)
	}
	return false, nil
}

func (r *StringResp) Parse() (string, error) {
	if r == nil {
		return "", fmt.Errorf("response is nil")
	}
	if IsNoDataFound(r.Message) { // 没有找到数据
		return "", nil
	}
	if r.Status == 0 {
		return "", fmt.Errorf("status not ok: %s, result: %s", r.Message, r.Result)
	}
	return r.Result, nil
}
