package etherscan

import (
	"context"
	"encoding/json"
	"fmt"
)

var gasTrackerModuleParams = map[string]string{
	"module": "gastracker",
}

// GetGasOracle Get current gas price recommendations.
func (c *client) GetGasOracle(ctx context.Context, chainID uint64) (resp *GasOracleResp, err error) {
	params := CopyMap(gasTrackerModuleParams)
	params["action"] = "gasoracle"
	params["chainId"] = fmt.Sprintf("%d", chainID)

	_, err = c.resty.R().SetContext(ctx).SetQueryParams(params).SetError(&resp).SetResult(&resp).Get("")
	return
}

// ----------------------------	REQUEST	------------------------------------

// ----------------------------	RESPONSE------------------------------------

type GasOracleResp BaseResp

type GasOracle struct {
	LastBlock       string `json:"LastBlock"`
	SafeGasPrice    string `json:"SafeGasPrice"`
	ProposeGasPrice string `json:"ProposeGasPrice"`
	FastGasPrice    string `json:"FastGasPrice"`
	SuggestBaseFee  string `json:"suggestBaseFee"`
	GasUsedRatio    string `json:"gasUsedRatio"`
}

func (r *GasOracleResp) GetData() (*GasOracle, error) {
	noData, err := (*BaseResp)(r).Parse()
	if err != nil {
		return nil, err
	}
	if noData {
		return nil, nil
	}

	var gasOracle GasOracle
	err = json.Unmarshal(r.Result, &gasOracle)
	if err != nil {
		return nil, fmt.Errorf("failed to parse result: %v", err)
	}
	return &gasOracle, nil
}
