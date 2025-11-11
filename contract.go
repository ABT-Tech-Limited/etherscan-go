package etherscan

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

var contractModuleParams = map[string]string{
	"module": "contract",
}

// GetContractABI Returns the Contract Application Binary Interface ( ABI ) of a verified smart contract.
func (c *client) GetContractABI(req GetContractABIReq) (resp *StringResp, err error) {
	params := CopyMap(contractModuleParams)
	params["action"] = "getabi"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// GetContractSourceCode Returns the Solidity source code of a verified smart contract.
func (c *client) GetContractSourceCode(req GetContractSourceCodeReq) (resp *ContractSourcecodeResp, err error) {
	params := CopyMap(contractModuleParams)
	params["action"] = "getsourcecode"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return

}

// GetContractCreatorTxInfo Returns a contract's deployer address and transaction hash it was created, up to 5 at a time.
func (c *client) GetContractCreatorTxInfo(req GetContractCreatorTxInfoReq) (resp *ContractCreatorTxInfoResp, err error) {
	params := CopyMap(contractModuleParams)
	params["action"] = "getcontractcreation"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	params["contractaddresses"] = strings.Join(req.Addresses, ",")
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// VerifySourceCode Submits a contract source code to an Etherscan-like explorer for verification.
func (c *client) VerifySourceCode(req VerifySourceCodeReq) (resp *VerifySourceCodeResp, err error) {
	params := CopyMap(contractModuleParams)
	params["action"] = "verifysourcecode"
	params["chainid"] = strconv.FormatUint(req.ChainID, 10)
	_, err = c.resty.R().SetQueryParams(params).SetFormData(StructToMap(req)).SetError(&resp).SetResult(&resp).Post("")
	return
}

// CheckVerifyStatus Returns the success or error status of a contract verification request.
func (c *client) CheckVerifyStatus(req CheckVerifyStatusReq) (resp *StringResp, err error) {
	params := CopyMap(contractModuleParams)
	params["action"] = "checkverifystatus"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// ----------------------------	REQUEST	------------------------------------

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

// ----------------------------	RESPONSE------------------------------------

type ContractSourcecodeResp BaseResp

func (r *ContractSourcecodeResp) GetData() ([]ContractSourceCode, error) {
	noData, err := (*BaseResp)(r).Parse()
	if err != nil {
		return nil, err
	}
	if noData {
		return nil, nil
	}

	var codes []ContractSourceCode
	err = json.Unmarshal(r.Result, &codes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return codes, nil
}

type ContractCreatorTxInfoResp BaseResp

func (r *ContractCreatorTxInfoResp) GetData() ([]ContractCreatorTxInfo, error) {
	noData, err := (*BaseResp)(r).Parse()
	if err != nil {
		return nil, err
	}
	if noData {
		return nil, nil
	}

	var infos []ContractCreatorTxInfo
	err = json.Unmarshal(r.Result, &infos)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return infos, nil
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

type ContractCreatorTxInfo struct {
	ContractAddress  string `json:"contractAddress"`
	ContractCreator  string `json:"contractCreator"`
	TxHash           string `json:"txHash"`
	BlockNumber      string `json:"blockNumber"`
	Timestamp        string `json:"timeStamp"`
	ContractFactory  string `json:"contractFactory"`
	CreationBytecode string `json:"creationBytecode"`
}

type VerifySourceCodeResp StringResp

func (r *VerifySourceCodeResp) GetData() (string, error) {
	result, err := (*StringResp)(r).Parse()
	if err != nil {
		return "", err
	}

	if result == ContractCodeAlreadyVerified {
		return "", nil
	}
	return "", fmt.Errorf("unexpected result: %s", result)
}
