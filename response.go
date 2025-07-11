package etherscan

type BaseResp struct {
	Status  int    `json:"status,string"` // 1 for good, 0 for error
	Message string `json:"message"`       // OK for good, other words when Status equals 0
}

// Module: Log

type LogResp struct {
	BaseResp
	Result []Log `json:"result"`
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
