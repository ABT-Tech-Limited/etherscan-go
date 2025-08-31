package etherscan

import (
	"strings"
)

var contractModuleParams = map[string]string{
	"module": "contract",
}

// GetContractABI Returns the Contract Application Binary Interface ( ABI ) of a verified smart contract.
func (c *Client) GetContractABI(req GetContractABIReq) (resp *ContractABIResp, err error) {
	params := contractModuleParams
	params["action"] = "getabi"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// GetContractSourceCode Returns the Solidity source code of a verified smart contract.
func (c *Client) GetContractSourceCode(req GetContractSourceCodeReq) (resp *ContractSourcecodeResp, err error) {
	params := contractModuleParams
	params["action"] = "getsourcecode"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return

}

// GetContractCreatorTxInfo Returns a contract's deployer address and transaction hash it was created, up to 5 at a time.
func (c *Client) GetContractCreatorTxInfo(req GetContractCreatorTxInfoReq) (resp *ContractCreatorTxInfoResp, err error) {
	params := contractModuleParams
	params["action"] = "getcontractcreation"
	for k, v := range StructToMap(req) {
		params[k] = v
	}
	params["contractaddresses"] = strings.Join(req.Addresses, ",")
	_, err = c.resty.R().SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}
