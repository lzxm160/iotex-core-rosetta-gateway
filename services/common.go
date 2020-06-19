package services

import "github.com/coinbase/rosetta-sdk-go/types"

//
//import (
//	"context"
//
//	"github.com/coinbase/rosetta-sdk-go/types"
//
//	oc "github.com/oasisprotocol/oasis-core-rosetta-gateway/oasis-client"
//)
//
//// OasisBlockchainName is the name of the Oasis blockchain.
//const OasisBlockchainName = "Oasis"
//
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
//// ValidateNetworkIdentifier validates the network identifier.
//func ValidateNetworkIdentifier(ctx context.Context, oc oc.OasisClient, ni *types.NetworkIdentifier) *types.Error {
//	if ni != nil {
//		if ni.Blockchain != OasisBlockchainName {
//			return ErrInvalidBlockchain
//		}
//		if ni.SubNetworkIdentifier != nil {
//			return ErrInvalidSubnetwork
//		}
//		chainID, err := GetChainID(ctx, oc)
//		if err != nil {
//			return err
//		}
//		if ni.Network != chainID {
//			return ErrInvalidNetwork
//		}
//	} else {
//		return ErrMissingNID
//	}
//	return nil
//}
