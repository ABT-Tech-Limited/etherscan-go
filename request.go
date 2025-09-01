package etherscan

// Module: Contract

type GetContractABIReq struct {
	ChainID uint64 `json:"chainid"`
	Address string `json:"address"` // the string representing the address of the contract
}

type GetContractSourceCodeReq struct {
	ChainID uint64 `json:"chainid"`
	Address string `json:"address"` // the string representing the address of the contract
}

type GetContractCreatorTxInfoReq struct {
	ChainID   uint64   `json:"chainid"`
	Addresses []string `json:"contractaddresses"` // the array representing the addresses of the contract
}

type VerifySourceCodeReq struct {
	ChainID              uint64  `json:"chainid"`
	CodeFormat           string  `json:"codeformat"`            // single file, use solidity-single-file JSON file ( recommended ), use solidity-standard-json-input
	SourceCode           string  `json:"sourceCode"`            // the Solidity source code
	ContractAddress      string  `json:"contractaddress"`       // the address your contract is deployed at
	ContractName         string  `json:"contractname"`          // the name of your contract, such as contracts/Verified.sol:Verified
	CompilerVersion      string  `json:"compilerversion"`       // compiler version used, such as v0.8.24+commit.e11b9ed9
	ConstructorArguments *string `json:"constructorArguements"` // optional, include if your contract uses constructor arguments
	CompilerMode         *string `json:"compilermode"`          // for ZK Stack, set to solc/zksync
	ZkSolcVersion        *string `json:"zksolcVersion"`         // for ZK Stack, zkSolc version used, such as v1.3.14
}

type CheckVerifyStatusReq struct {
	ChainID uint64 `json:"chainid"`
	GUID    string `json:"guid"` // the guid returned from the verify API
}

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
