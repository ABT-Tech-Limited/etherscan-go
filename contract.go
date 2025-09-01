package etherscan

import (
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
