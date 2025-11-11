package etherscan

import (
	"encoding/json"
	"fmt"
)

var accountModuleParams = map[string]string{
	"module": "account",
}

func (c *client) GetNormalTransactions(req GetNormalTransactionsReq) (resp *TransactionListResp, err error) {
	params := CopyMap(accountModuleParams)
	params["action"] = "txlist"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// ----------------------------	REQUEST	------------------------------------

type GetNormalTransactionsReq struct {
	ChainID    uint64 `json:"chainid"`
	Address    string `json:"address"`
	Page       uint64 `json:"page"`
	Offset     uint64 `json:"offset"`
	StartBlock uint64 `json:"startblock,omitempty"`
	EndBlock   uint64 `json:"endblock,omitempty"`
	Sort       string `json:"sort,omitempty"` // asc or desc
}

// ----------------------------	RESPONSE------------------------------------

type TransactionListResp BaseResp

func (r *TransactionListResp) GetData() ([]Transaction, error) {
	noData, err := (*BaseResp)(r).Parse()
	if err != nil {
		return nil, err
	}
	if noData {
		return nil, nil
	}

	var txs []Transaction
	err = json.Unmarshal(r.Result, &txs)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return txs, nil
}

type Transaction struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	TransactionIndex  string `json:"transactionIndex"`
	From              string `json:"from"`
	To                string `json:"to"`
	Value             string `json:"value"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	IsError           string `json:"isError"`
	TxreceiptStatus   string `json:"txreceipt_status"`
	Input             string `json:"input"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	GasUsed           string `json:"gasUsed"`
	Confirmations     string `json:"confirmations"`
	MethodID          string `json:"methodId"`
	FunctionName      string `json:"functionName"`
}
