package services

import (
	"context"
	"fmt"

	ic "github.com/iotexproject/iotex-core-rosetta-gateway/iotex-client"

	"github.com/coinbase/rosetta-sdk-go/types"
)

//
//import (
//	"context"
//
//	"github.com/coinbase/rosetta-sdk-go/types"
//
//	oc "github.com/oasisprotocol/oasis-core-rosetta-gateway/oasis-client"
//)
//
// BlockchainName is the name of the IoTex blockchain.
const (
	BlockchainName = "IoTex"
	chainID        = "mainnet"
)

// IoTexCurrency is the currency used on the IoTex blockchain.
var IoTexCurrency = &types.Currency{
	Symbol:   "IoTex",
	Decimals: 18,
}

//
//// GetChainID returns the chain ID.
//func GetChainID(ctx context.Context, oc oc.OasisClient) (string, *types.Error) {
//	chainID, err := oc.GetChainID(ctx)
//	if err != nil {
//		return "", ErrUnableToGetChainID
//	}
//	return chainID, nil
//}
//
// ValidateNetworkIdentifier validates the network identifier.
func ValidateNetworkIdentifier(ctx context.Context, c ic.IoTexClient, ni *types.NetworkIdentifier) *types.Error {
	if ni != nil {
		if ni.Blockchain != BlockchainName {
			return ErrInvalidBlockchain
		}
		if ni.SubNetworkIdentifier != nil {
			fmt.Println("ni.SubNetworkIdentifier != nil")
			return ErrInvalidSubnetwork
		}
		if ni.Network != chainID {
			return ErrInvalidNetwork
		}
	} else {
		return ErrMissingNID
	}
	return nil
}
