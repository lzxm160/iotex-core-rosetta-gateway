package iotex_client

import (
	//"encoding/hex"
	//"fmt"
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc/credentials"

	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-proto/golang/iotexapi"

	"github.com/iotexproject/iotex-antenna-go/v2/iotex"
	//"github.com/trustwallet/blockatlas/pkg/logger"
	//"os"

	//"context"
	//"encoding/hex"
	//"encoding/json"
	//"fmt"
	//"os"
	//"sync"
	//
	"github.com/coinbase/rosetta-sdk-go/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	//cmnGrpc "github.com/oasisprotocol/oasis-core/go/common/grpc"
	//"github.com/oasisprotocol/oasis-core/go/common/logging"
	//consensus "github.com/oasisprotocol/oasis-core/go/consensus/api"
	//"github.com/oasisprotocol/oasis-core/go/consensus/api/transaction"
	//control "github.com/oasisprotocol/oasis-core/go/control/api"
	//staking "github.com/oasisprotocol/oasis-core/go/staking/api"

	"sync"
)

var IoTexCurrency = &types.Currency{
	Symbol:   "IoTex",
	Decimals: 18,
}

// IoTexClient is the IoTex blockchain client interface.
type IoTexClient interface {
	// GetChainID returns the network chain context, derived from the
	// genesis document.
	GetChainID(ctx context.Context) (string, error)

	// GetBlock returns the Oasis block at given height.
	GetBlock(ctx context.Context, height int64) (*IoTexBlock, error)

	// GetLatestBlock returns latest Oasis block.
	GetLatestBlock(ctx context.Context) (*IoTexBlock, error)

	// GetGenesisBlock returns the Oasis genesis block.
	GetGenesisBlock(ctx context.Context) (*IoTexBlock, error)

	// GetAccount returns the Oasis staking account for given owner address
	// at given height.
	GetAccount(ctx context.Context, height int64, owner string) (*Account, error)
	//
	//// GetStakingEvents returns Oasis staking events at given height.
	//GetStakingEvents(ctx context.Context, height int64) ([]staking.Event, error)
	//
	// SubmitTx submits the given JSON-encoded transaction to the node.
	SubmitTx(ctx context.Context, tx *iotextypes.Action) (txid string, err error)
	//
	//// GetNextNonce returns the nonce that should be used when signing the
	//// next transaction for the given account address at given height.
	//GetNextNonce(ctx context.Context, addr staking.Address, height int64) (uint64, error)
	//
	// GetStatus returns the status overview of the node.
	GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error)
	GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error)
	GetTransactions(ctx context.Context, height int64) ([]*types.Transaction, error)
}

// IoTexBlock is the IoTex blockchain's block.
type IoTexBlock struct {
	Height       int64  // Block height.
	Hash         string // Block hash.
	Timestamp    int64  // UNIX time, converted to milliseconds.
	ParentHeight int64  // Height of parent block.
	ParentHash   string // Hash of parent block.
}

type Account struct {
	Nonce   uint64
	Balance string
}

// grpcIoTexClient is an implementation of IoTexClient using gRPC.
type grpcIoTexClient struct {
	sync.RWMutex

	endpoint string
	grpcConn *grpc.ClientConn
}

// NewIoTexClient returns an implementation of IoTexClient
func NewIoTexClient(grpcAddr string) (cli IoTexClient, err error) {
	grpc, err := newDefaultGRPCConn(grpcAddr)
	if err != nil {
		return
	}
	cli = &grpcIoTexClient{endpoint: grpcAddr, grpcConn: grpc}
	return
}

func (c *grpcIoTexClient) reconnect(ctx context.Context) {
	c.Lock()
	defer c.Unlock()

	// Check if the existing connection is good.
	if c.grpcConn != nil && c.grpcConn.GetState() != connectivity.Shutdown {
		return
	} else {
		// Connection needs to be re-established.
		c.grpcConn = nil
	}
	c.grpcConn, _ = iotex.NewDefaultGRPCConn(c.endpoint)
}

func (c *grpcIoTexClient) GetChainID(ctx context.Context) (string, error) {
	c.reconnect(ctx)
	c.Lock()
	defer c.Unlock()
	return "1", nil
	//client := consensus.NewConsensusClient(conn)
	//genesis, err := client.GetGenesisDocument(ctx)
	//if err != nil {
	//	logger.Debug("GetChainID: failed to get genesis document", "err", err)
	//	return "", err
	//}
	//oc.chainID = genesis.ChainContext()
	//return oc.chainID, nil
}

func (c *grpcIoTexClient) GetBlock(ctx context.Context, height int64) (ret *IoTexBlock, err error) {
	c.reconnect(ctx)
	var parentHeight uint64
	if height <= 1 {
		parentHeight = 1
	} else {
		parentHeight = uint64(height) - 1
	}
	fmt.Println(parentHeight, height)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	count := uint64(2)
	if parentHeight == uint64(height) {
		count = 1
	}
	request := &iotexapi.GetBlockMetasRequest{
		Lookup: &iotexapi.GetBlockMetasRequest_ByIndex{
			ByIndex: &iotexapi.GetBlockMetasByIndexRequest{
				Start: parentHeight,
				Count: count,
			},
		},
	}
	resp, err := client.GetBlockMetas(ctx, request)
	if err != nil {
		return nil, err
	}
	if len(resp.BlkMetas) == 0 {
		return nil, errors.New("not found")
	}
	var blk, parentBlk *iotextypes.BlockMeta
	if len(resp.BlkMetas) == 2 {
		blk = resp.BlkMetas[1]
		parentBlk = resp.BlkMetas[0]
		fmt.Println("len(resp.BlkMetas) == 2")
	} else {
		blk = resp.BlkMetas[0]
		parentBlk = resp.BlkMetas[0]
	}
	ret = &IoTexBlock{
		Height:       int64(blk.Height),
		Hash:         blk.Hash,
		Timestamp:    (blk.Timestamp.Seconds*1e9 + int64(blk.Timestamp.Nanos)) / 1e6, // ms
		ParentHeight: int64(parentHeight),
		ParentHash:   parentBlk.Hash,
	}
	return
}

func (c *grpcIoTexClient) GetLatestBlock(ctx context.Context) (*IoTexBlock, error) {
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	res, err := client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
	if err != nil {
		return nil, err
	}
	return c.GetBlock(ctx, int64(res.ChainMeta.Height))
}

func (c *grpcIoTexClient) GetGenesisBlock(ctx context.Context) (*IoTexBlock, error) {
	return c.GetBlock(ctx, 1)
}

func (c *grpcIoTexClient) GetAccount(ctx context.Context, height int64, owner string) (ret *Account, err error) {
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	request := &iotexapi.GetAccountRequest{Address: owner}
	resp, err := client.GetAccount(ctx, request)
	if err != nil {
		return nil, err
	}
	ret = &Account{
		Nonce:   resp.AccountMeta.Nonce,
		Balance: resp.AccountMeta.Balance,
	}
	return
}

func (c *grpcIoTexClient) GetTransactions(ctx context.Context, height int64) (ret []*types.Transaction, err error) {
	blk, err := c.GetBlock(ctx, height)
	if err != nil {
		return
	}
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	fmt.Println("before client.GetActions")
	limit := uint64(1000)
	var actionInfo []*iotexapi.ActionInfo
	for i := uint64(0); ; i++ {
		request := &iotexapi.GetActionsRequest{
			Lookup: &iotexapi.GetActionsRequest_ByBlk{
				ByBlk: &iotexapi.GetActionsByBlockRequest{
					BlkHash: blk.Hash,
					Start:   i * limit,
					Count:   limit,
				},
			},
		}
		res, err := client.GetActions(context.Background(), request)
		if err != nil {
			break
		}
		actionInfo = append(actionInfo, res.ActionInfo...)
	}

	fmt.Println("after client.GetActions")
	ret = make([]*types.Transaction, 0)
	for _, act := range actionInfo {
		transfer := act.GetAction().GetCore().GetTransfer()
		if transfer == nil {
			continue
		}
		requestGetReceipt := &iotexapi.GetReceiptByActionRequest{ActionHash: act.GetActHash()}
		responseReceipt, err := client.GetReceiptByAction(ctx, requestGetReceipt)
		if err != nil {
			continue
		}
		status := "succeed"
		if responseReceipt.GetReceiptInfo().GetReceipt().GetStatus() != 1 {
			status = "fail"
		}
		oper := []*types.Operation{
			&types.Operation{
				OperationIdentifier: &types.OperationIdentifier{
					Index:        int64(act.GetAction().GetCore().GetNonce()),
					NetworkIndex: nil,
				},
				RelatedOperations: nil,
				Type:              "transfer",
				Status:            status,
				Account: &types.AccountIdentifier{
					Address:    act.Sender,
					SubAccount: nil,
					Metadata:   nil,
				},
				Amount: &types.Amount{
					Value:    transfer.Amount,
					Currency: IoTexCurrency,
					Metadata: nil,
				},
				Metadata: nil,
			},
		}

		ret = append(ret, &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{
				act.ActHash,
			},
			Operations: oper,
			Metadata:   nil,
		})
	}
	return
}

func (c *grpcIoTexClient) SubmitTx(ctx context.Context, tx *iotextypes.Action) (txid string, err error) {
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	ret, err := client.SendAction(ctx, &iotexapi.SendActionRequest{Action: tx})
	if err != nil {
		return
	}
	txid = ret.ActionHash
	return
}

func (c *grpcIoTexClient) GetStatus(ctx context.Context) (*iotexapi.GetChainMetaResponse, error) {
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	return client.GetChainMeta(context.Background(), &iotexapi.GetChainMetaRequest{})
}

func (c *grpcIoTexClient) GetVersion(ctx context.Context) (*iotexapi.GetServerMetaResponse, error) {
	c.reconnect(ctx)
	client := iotexapi.NewAPIServiceClient(c.grpcConn)
	return client.GetServerMeta(context.Background(), nil)
}

// newDefaultGRPCConn creates a default grpc connection. With tls and retry.
func newDefaultGRPCConn(endpoint string) (*grpc.ClientConn, error) {
	opts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffLinear(10 * time.Second)),
		grpc_retry.WithMax(3),
	}
	return grpc.Dial(endpoint,
		grpc.WithStreamInterceptor(grpc_retry.StreamClientInterceptor(opts...)),
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(opts...)),
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
}
