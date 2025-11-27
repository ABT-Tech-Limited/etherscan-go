package etherscan

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetContractABI(t *testing.T) {
	rsp, err := api.GetContractABI(context.Background(), GetContractABIReq{
		ChainID: chainIDETH,
		Address: "0xBB9bc244D798123fDe783fCc1C72d3Bb8C189413",
	})
	assert.Nil(t, err)
	assert.Equal(t, 1, rsp.Status)
	assert.NotEmpty(t, rsp.Result)
}

func TestClient_GetContractSourceCode(t *testing.T) {
	rsp, err := api.GetContractSourceCode(context.Background(), GetContractSourceCodeReq{
		ChainID: chainIDETH,
		Address: "0xBB9bc244D798123fDe783fCc1C72d3Bb8C189413",
	})
	assert.Nil(t, err)
	code, err := rsp.GetData()
	assert.Nil(t, err)
	assert.Len(t, code, 1)
	assert.Equal(t, "DAO", code[0].ContractName)
}

func TestClient_GetContractCreatorTxInfo(t *testing.T) {
	rsp, err := api.GetContractCreatorTxInfo(context.Background(), GetContractCreatorTxInfoReq{
		ChainID: chainIDETH,
		Addresses: []string{
			"0xbb9bc244d798123fde783fcc1c72d3bb8c189413",
			"0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45",
			"0xb83c27805aaca5c7082eb45c868d955cf04c337f",
			"0xe4462eb568e2dfbb5b0ca2d3dbb1a35c9aa98aad",
			"0xdac17f958d2ee523a2206206994597c13d831ec7",
		},
	})
	assert.Nil(t, err)
	info, err := rsp.GetData()
	assert.Nil(t, err)
	assert.Len(t, info, 5)
}

func TestClient_VerifySourceCode(t *testing.T) {
	rsp, err := api.VerifySourceCode(context.Background(), VerifySourceCodeReq{
		ChainID:              chainIDETH,
		CodeFormat:           "",
		SourceCode:           "",
		ContractAddress:      "",
		ContractName:         "",
		CompilerVersion:      "",
		ConstructorArguments: nil,
		CompilerMode:         nil,
		ZkSolcVersion:        nil,
	})
	assert.Nil(t, err)
	assert.NotNil(t, rsp)
}

func TestClient_CheckVerifyStatus(t *testing.T) {
	rsp, err := api.CheckVerifyStatus(context.Background(), CheckVerifyStatusReq{
		ChainID: chainIDETH,
		GUID:    "x3ryqcqr1zdknhfhkimqmizlcqpxncqc6nrvp3pgrcpfsqedqi",
	})
	assert.Nil(t, err)
	assert.Equal(t, "Pass - Verified", rsp.Result)
}
