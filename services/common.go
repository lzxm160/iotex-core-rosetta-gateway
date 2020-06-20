package services

import (
	"context"
	"fmt"

	"github.com/coinbase/rosetta-sdk-go/types"

	ic "github.com/iotexproject/iotex-core-rosetta-gateway/iotex-client"
)

// IoTexCurrency is the currency used on the IoTex blockchain.
var IoTexCurrency = &types.Currency{
	Symbol:   "IoTex",
	Decimals: 18,
}

// ValidateNetworkIdentifier validates the network identifier.
func ValidateNetworkIdentifier(ctx context.Context, c ic.IoTexClient, ni *types.NetworkIdentifier) *types.Error {
	if ni != nil {
		fmt.Println("ni != nil")
		cfg := c.GetConfig()
		if ni.Blockchain != cfg.Network_identifier.Blockchain {
			return ErrInvalidBlockchain
		}
		if ni.SubNetworkIdentifier != nil {
			fmt.Println("ni.SubNetworkIdentifier != nil")
			return ErrInvalidSubnetwork
		}
		fmt.Println("ni.Network chainID")
		if ni.Network != cfg.Network_identifier.Network {
			return ErrInvalidNetwork
		}
	} else {
		return ErrMissingNID
	}
	return nil
}
