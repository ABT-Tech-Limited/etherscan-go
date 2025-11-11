package etherscan

// Module: Account

// Module: Contract

// Module: Log

type GetEventLogsByAddressReq struct {
	ChainID   uint64 `json:"chainid"`
	Address   string `json:"address"`   // the string representing the address to check for logs
	FromBlock uint64 `json:"fromBlock"` // the integer block number to start searching for logs eg. 12878196
	ToBlock   uint64 `json:"toBlock"`   // the integer block number to stop searching for logs eg. 12879196
	Page      *int   `json:"page"`      // the integer page number, if pagination is enabled
	Offset    *int   `json:"offset"`    // the number of transactions displayed per page, limited to 1000 records per query, use the page parameter for subsequent records
}

type GetEventLogsByTopicsReq struct {
	ChainID   uint64            `json:"chainid"`
	FromBlock uint64            `json:"fromBlock"` // the integer block number to start searching for logs eg. 12878196
	ToBlock   uint64            `json:"toBlock"`   // the integer block number to stop searching for logs eg. 12879196
	Topics    map[string]string `json:"topics"`    // topic & topicOperator
	Page      *int              `json:"page"`      // the integer page number, if pagination is enabled
	Offset    *int              `json:"offset"`    // the number of transactions displayed per page, limited to 1000 records per query, use the page parameter for subsequent records
}

type GetEventLogsByAddressFilterByTopicsReq struct {
	ChainID   uint64            `json:"chainid"`
	Address   string            `json:"address"`   // the string representing the address to check for logs
	FromBlock uint64            `json:"fromBlock"` // the integer block number to start searching for logs eg. 12878196
	ToBlock   uint64            `json:"toBlock"`   // the integer block number to stop searching for logs eg. 12879196
	Topics    map[string]string `json:"topics"`    // topic & topicOperator
	Page      *int              `json:"page"`      // the integer page number, if pagination is enabled
	Offset    *int              `json:"offset"`    // the number of transactions displayed per page, limited to 1000 records per query, use the page parameter for subsequent records
}

// Topics
// topic0/topic1/topic2/topic3
// topic0_1_opr=and topic0_2_opr=or ...
