// Copyright (c) 2020 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package inject

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-antenna-go/v2/account"
	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

const (
	sender     = "io1ph0u2psnd7muq5xv9623rmxdsxc4uapxhzpg02"
	to         = "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
	privateKey = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	//endpoint         = "127.0.0.1:14014"
	endpoint         = "api.testnet.iotex.one:80"
	IoTeXDID_address = "io14gqv7s4dkfhdgssq4l7sedjv68kv70hl4x5j0u"
)

var (
	gasPrice = big.NewInt(0).SetUint64(1e12)
	gasLimit = uint64(10000000)
)

func TestInjectTransfer(t *testing.T) {
	fmt.Println("inject transfer")
	require := require.New(t)
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	to, err := address.FromString(to)
	require.NoError(err)
	hash, err := c.Transfer(to, big.NewInt(0).SetUint64(1000)).SetGasPrice(gasPrice).SetGasLimit(gasLimit).Call(context.
		Background())
	require.NoError(err)
	require.NotNil(hash)
	checkHash(hex.EncodeToString(hash[:]), t)
}

func TestDidDeployContract(t *testing.T) {
	//require := require.New(t)
	//conn, err := iotex.NewDefaultGRPCConn(endpoint)
	//require.NoError(err)
	//defer conn.Close()
	//
	//acc, err := account.HexStringToAccount(privateKey)
	//require.NoError(err)
	//c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	//
	//data, err := hex.DecodeString(abibin.AddressBasedDIDManagerWithAgentEnabledBin[2:])
	//require.NoError(err)
	//abi, err := abi.JSON(strings.NewReader(abibin.AddressBasedDIDManagerWithAgentEnabledABI))
	//require.NoError(err)
	//hash, err := c.DeployContract(data).SetArgs(abi, []byte("did:io:"), common.Address{0}).SetGasPrice(big.NewInt(int64(unit.Qev))).SetGasLimit(10000000).Call(context.Background())
	//require.NoError(err)
	//require.NotNil(hash)
	//fmt.Println("hash", hex.EncodeToString(hash[:]))
	//time.Sleep(20 * time.Second)
	//receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	//require.NoError(err)
	//contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	//fmt.Println("Status:", receiptResponse.GetReceiptInfo().GetReceipt().Status)
	//fmt.Println("Contract Address:", contractAddress)
}

func checkHash(h string, t *testing.T) {
	fmt.Println("check hash:", h)
	require := require.New(t)
	time.Sleep(20 * time.Second)
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	require.NoError(err)
	defer conn.Close()
	ha, err := hash.HexStringToHash256(h)
	require.NoError(err)
	c := iotex.NewReadOnlyClient(iotexapi.NewAPIServiceClient(conn))
	receiptResponse, err := c.GetReceipt(ha).Call(context.Background())
	s := receiptResponse.GetReceiptInfo().GetReceipt().GetStatus()
	fmt.Println("status:", s)
}
