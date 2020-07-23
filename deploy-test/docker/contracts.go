package docker

import (
	ethutils "github.com/ChainSafe/ChainBridge/shared/ethereum"
	"github.com/ethereum/go-ethereum/common"
)

var TestContracts = ethutils.DeployedContracts{
	BridgeAddress:         common.HexToAddress("0x62877dDCd49aD22f5eDfc6ac108e9a4b5D2bD88B"),
	ERC20HandlerAddress:   common.HexToAddress("0x3167776db165D8eA0f51790CA2bbf44Db5105ADF"),
	ERC721HandlerAddress:  common.HexToAddress("0x3f709398808af36ADBA86ACC617FeB7F5B7B193E"),
	GenericHandlerAddress: common.HexToAddress("0x2B6Ab4b880A45a07d83Cf4d664Df4Ab85705Bc07"),
}