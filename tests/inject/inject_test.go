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
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	privateKey = "414efa99dfac6f4095d6954713fb0085268d400d6a05a8ae8a69b5b1c10b4bed"
	to         = "io1vdtfpzkwpyngzvx7u2mauepnzja7kd5rryp0sg"
	receipt    = "io1mflp9m6hcgm2qcghchsdqj3z3eccrnekx9p0ms"
	endpoint   = "127.0.0.1:14014"
	//endpoint         = "api.testnet.iotex.one:80"
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
	hash, err := c.Transfer(to, big.NewInt(0).SetUint64(1000)).SetGasPrice(gasPrice).SetGasLimit(gasLimit).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	checkHash(hex.EncodeToString(hash[:]), t)
}

func TestMultisend(t *testing.T) {
	require := require.New(t)
	contract := deployContract(t)
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	abi, err := abi.JSON(strings.NewReader(MultisendABI))
	require.NoError(err)
	contractAddr, err := address.FromString(contract)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)
	r1, err := address.FromString(to)
	require.NoError(err)
	r2, err := address.FromString(receipt)
	require.NoError(err)
	r1ethAddress := common.HexToAddress(hex.EncodeToString(r1.Bytes()))
	r2ethAddress := common.HexToAddress(hex.EncodeToString(r2.Bytes()))
	hash, err := c.Contract(contractAddr, abi).Execute("multiSend", []common.Address{r1ethAddress, r2ethAddress}, []*big.Int{big.NewInt(1), big.NewInt(2)}, "").SetGasPrice(gasPrice).SetGasLimit(gasLimit).SetAmount(big.NewInt(3)).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	checkHash(hex.EncodeToString(hash[:]), t)
}

func deployContract(t *testing.T) string {
	require := require.New(t)
	conn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	require.NoError(err)
	defer conn.Close()

	acc, err := account.HexStringToAccount(privateKey)
	require.NoError(err)
	c := iotex.NewAuthedClient(iotexapi.NewAPIServiceClient(conn), acc)

	data, err := hex.DecodeString(MultisendBin[2:])
	require.NoError(err)

	hash, err := c.DeployContract(data).SetGasPrice(gasPrice).SetGasLimit(gasLimit).Call(context.Background())
	require.NoError(err)
	require.NotNil(hash)
	fmt.Println("hash", hex.EncodeToString(hash[:]))
	time.Sleep(15 * time.Second)
	receiptResponse, err := c.GetReceipt(hash).Call(context.Background())
	require.NoError(err)
	contractAddress := receiptResponse.GetReceiptInfo().GetReceipt().GetContractAddress()
	fmt.Println("Status:", receiptResponse.GetReceiptInfo().GetReceipt().Status)
	fmt.Println("Contract Address:", contractAddress)
	return contractAddress
}

func checkHash(h string, t *testing.T) {
	fmt.Println("check hash:", h)
	require := require.New(t)
	time.Sleep(10 * time.Second)
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
